#!/usr/bin/env bash
# anydoc-go 构建管线：
#   fetch: 拉取/刷新第三方仓库并应用本地 patch
#          URL 可用 ANYDOC_REPO / GOCCY_REPO 覆盖，版本可用 ANYDOC_REF / GOCCY_REF 覆盖
#   wasm:  cargo → target/<arch>/release/anydoc_cabi.wasm（依赖 third-party/anydoc）
#   cli:   goccy-wasm2go → core/ → go build → bin/anydoc
#   test:  go test ./...
#   build: wasm + cli（默认）
set -euo pipefail

ROOT="$(cd "$(dirname "$0")" && pwd)"
export RUSTUP_HOME="${RUSTUP_HOME:-${ROOT}/.rustup}"
export CARGO_HOME="${CARGO_HOME:-${ROOT}/.cargo}"
export CARGO_TARGET_DIR="${CARGO_TARGET_DIR:-${ROOT}/target}"
export ZIG_GLOBAL_CACHE_DIR="${ZIG_GLOBAL_CACHE_DIR:-${ROOT}/.zig-cache}"
export ZIG_LOCAL_CACHE_DIR="${ZIG_LOCAL_CACHE_DIR:-${ROOT}/libzig-cache}"
export GOCACHE="${GOCACHE:-${ROOT}/.go-cache}"
export GOPATH="${GOPATH:-${ROOT}/.gopath}"
export PATH="${CARGO_HOME}/bin:${ROOT}/.zig/bin:${ROOT}/bin:${PATH}"
# 仅本机用 zig cc 作宿主链接器（无系统 cc；.zig 见 README）。
# 没有 .zig 的环境（如 CI runner，自带 gcc）直接走系统 cc。
if [ -x "${ROOT}/.zig/bin/zig" ]; then
    export CC="${CC:-zig cc}"
fi

ANYDOC_DIR="${ROOT}/third-party/anydoc"
GOCCY_DIR="${ROOT}/third-party/goccy-wasm2go"
ANYDOC_REPO="${ANYDOC_REPO:-https://github.com/firecrawl/anydoc}"
GOCCY_REPO="${GOCCY_REPO:-https://github.com/goccy/wasm2go}"

FETCH_REF="${FETCH_REF:-}"           # 同时覆盖两个第三方仓库（兼容旧用法）
ANYDOC_REF="${ANYDOC_REF:-}"         # 空值：取最新 semver tag
GOCCY_REF="${GOCCY_REF:-f30ec292fd4ea1737263c30ed97157e4593796db}"
WASM="${CARGO_TARGET_DIR}/wasm32-unknown-unknown/release/anydoc_cabi.wasm"

# 语义化版本排序的最新 tag（v0.1.9 > v0.1.10 > v0.2.0）。
latest_tag() {
    git ls-remote --tags --refs --sort=-v:refname "$1" 'refs/tags/*' 2>/dev/null \
        | head -1 | awk '{print $NF}' | sed 's|refs/tags/||'
}

fetch() {
    # 第三方仓库位于 anydoc-go 自己的 third-party/ 下（.gitignore 排除，
    # 不随本仓库提交）。anydoc 默认跟最新发布 tag；goccy 固定到已验证
    # revision，确保生成器版本稳定。
    fetch_one anydoc "${ANYDOC_REPO}" "${ANYDOC_DIR}" "${ANYDOC_REF}"
    fetch_one goccy  "${GOCCY_REPO}"  "${GOCCY_DIR}"  "${GOCCY_REF}"
    require_repos
    apply_anydoc_patch
    echo "third-party repos ready:"
    git -C "${ANYDOC_DIR}" rev-parse HEAD
    git -C "${GOCCY_DIR}" rev-parse HEAD
}

fetch_one() {
    local name="$1" url="$2" dir="$3" requested_ref="${4:-}"
    local ref="${FETCH_REF:-${requested_ref}}"
    if [ -z "${ref}" ]; then
        if [ "${name}" = "anydoc" ]; then
            ref="$(latest_tag "${url}")" || true
            [ -n "${ref}" ] || ref="main"   # 取 tag 失败时回退 main
        else
            ref="main"
        fi
        echo "${name}: target ${ref}"
    fi
    if [ ! -d "${dir}/.git" ]; then
        echo "cloning ${name} (${url}) → ${dir}"
        git clone -q "${url}" "${dir}"
    else
        echo "updating ${name} → ${dir}"
    fi

    # 先 fetch 精确目标，再以 detached HEAD 构建。这样 branch、tag、SHA
    # 都走同一条路径；目标发生变化而工作树有本地 patch 时直接失败，
    # 不会静默继续使用旧版本。
    git -C "${dir}" fetch -q --tags origin
    git -C "${dir}" fetch -q origin "${ref}"
    local target current
    target="$(git -C "${dir}" rev-parse FETCH_HEAD^{commit})"
    current="$(git -C "${dir}" rev-parse HEAD 2>/dev/null || true)"
    if [ "${current}" != "${target}" ]; then
        if ! git -C "${dir}" diff --quiet || ! git -C "${dir}" diff --cached --quiet; then
            echo "error: ${name} has local changes; cannot switch ${ref}" >&2
            exit 2
        fi
        git -C "${dir}" checkout -q --detach "${target}"
    fi
    echo "${name}: ${target}"
}

require_repos() {
    [ -d "${ANYDOC_DIR}/.git" ] || { echo "missing third-party/anydoc — run: $0 fetch" >&2; exit 2; }
    [ -d "${GOCCY_DIR}/.git" ] || { echo "missing third-party/goccy-wasm2go — run: $0 fetch" >&2; exit 2; }
}

# 应用本地 patch（第三方仓库不提交任何改动；CI fresh clone 后必须有同一
# patch 才能构建出 asset:// 占位）。
ANYDOC_PATCH="${ROOT}/patches/anydoc-asset-placeholder.patch"
apply_patch_once() {
    local repo="$1" patch="$2" name="$3"
    if git -C "${repo}" apply --check "${patch}" 2>/dev/null; then
        git -C "${repo}" apply "${patch}"
        echo "applied ${name} patch"
    elif git -C "${repo}" apply --reverse --check "${patch}" 2>/dev/null; then
        echo "${name} patch already applied"
    else
        echo "error: ${name} patch does not apply to $(git -C "${repo}" rev-parse HEAD)" >&2
        echo "       update ${patch} for this third-party revision" >&2
        exit 2
    fi
}

apply_anydoc_patch() {
    apply_patch_once "${ANYDOC_DIR}" "${ANYDOC_PATCH}" anydoc
}

wasm_step() {
    require_repos
    apply_anydoc_patch
    cd "${ROOT}/cabi"
    # 显式关 SIMD：goccy 支持 SIMD wasm，但本项目文档转换是标量/字符串/
    # XML/zlib 主导，实测 SIMD 反而慢 3-4 倍（见 README「工具链对比」）。
    RUSTFLAGS="-C target-feature=-simd128" cargo build --release --target wasm32-unknown-unknown
    cd "${ROOT}"
}

cli_step() {
    require_repos
    [ -f "${WASM}" ] || { echo "missing wasm — run: $0 wasm" >&2; exit 2; }
    # bin/ 不入库；每次在 patch 后重建，避免 ignored 的旧 translator 绕过补丁。
    echo "building goccy-wasm2go → ${ROOT}/bin/goccy-wasm2go"
    (cd "${GOCCY_DIR}" && go build -o "${ROOT}/bin/goccy-wasm2go" ./cmd/wasm2go)
    # goccy 生成纯 Go 多包 core（core.go + base/ + p0/ + p1/ + data.bin）。
    # -pure 禁止生成 asm bundle 和架构 build tags，保持跨平台纯 Go 构建；
    # 整目录进 core/，不需要 split_gen 拆文件。
    rm -rf "${ROOT}/core"
    "${ROOT}/bin/goccy-wasm2go" -pure -i "${WASM}" -pkg core -import github.com/xmdhs/anydoc-go/core \
        -out-dir "${ROOT}/core"
    cd "${ROOT}"
    go build -p 2 -trimpath -ldflags="-s -w" -o "${ROOT}/bin/anydoc" .
}

test_step() {
    cd "${ROOT}"
    go test -p 1 -trimpath -ldflags="-s -w" -v ./...
}

case "${1:-build}" in
    fetch) fetch ;;
    build)
        # wasm + cli 全流程（默认命令）
        wasm_step
        cli_step
        ;;
    wasm) wasm_step ;;
    cli) cli_step ;;
    test) test_step ;;
    *)
        echo "usage: $0 [fetch|wasm|cli|test|build]" >&2
        exit 2
        ;;
esac