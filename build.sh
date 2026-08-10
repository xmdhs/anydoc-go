#!/usr/bin/env bash
# anydoc-go 构建管线（goccy 分支）：
#   fetch: 拉取/刷新第三方仓库（third-party/anydoc、third-party/goccy-wasm2go），
#          URL 可用 ANYDOC_REPO / GOCCY2GO_REPO 覆盖
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

FETCH_REF="${FETCH_REF:-}"           # 显式指定分支/标签（默认见 fetch_one）
WASM="${CARGO_TARGET_DIR}/wasm32-unknown-unknown/release/anydoc_cabi.wasm"

# 语义化版本排序的最新 tag（v0.1.9 > v0.1.10 > v0.2.0）。
latest_tag() {
    git ls-remote --tags --refs --sort=-v:refname "$1" 'refs/tags/*' 2>/dev/null \
        | head -1 | awk '{print $NF}' | sed 's|refs/tags/||'
}

fetch() {
    # 第三方仓库位于 anydoc-go 自己的 third-party/ 下（.gitignore 排除，
    # 不随本仓库提交）。anydoc 只跟最新发布 tag（避免 main 的不稳定代码，
    # 也便于我们以 tag 为准复现构建）；goccy-wasm2go 跟 main。
    # fetch 只做 fetch/checkout/ff-only pull，有本地改动且冲突时不强推。
    fetch_one anydoc   "${ANYDOC_REPO}"   "${ANYDOC_DIR}"
    fetch_one goccy    "${GOCCY_REPO}"    "${GOCCY_DIR}"
    echo "third-party repos ready:" "${ANYDOC_DIR}" "${GOCCY_DIR}"
}

fetch_one() {
    local name="$1" url="$2" dir="$3"
    local ref="${FETCH_REF}"
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
        git -C "${dir}" fetch -q origin || true
        git -C "${dir}" fetch -q --tags origin || true
    fi
    # tag/branch 统一走 checkout；branch 再 ff-only pull 追平远端。
    git -C "${dir}" checkout -q "${ref}" \
        || git -C "${dir}" fetch -q origin tag "${ref}" \
        || echo "warn: checkout ${ref} 失败（本地改动冲突？）" >&2
    git -C "${dir}" pull -q --ff-only origin "${ref}" || true
}

require_repos() {
    [ -d "${ANYDOC_DIR}/.git" ] || { echo "missing third-party/anydoc — run: $0 fetch" >&2; exit 2; }
    [ -d "${GOCCY_DIR}/.git" ] || { echo "missing third-party/goccy-wasm2go — run: $0 fetch" >&2; exit 2; }
}

# 应用本地 patch（第三方仓库不提交任何改动；CI fresh clone 后必须有同一
# patch 才能构建出 asset:// 占位）。
ANYDOC_PATCH="${ROOT}/patches/anydoc-asset-placeholder.patch"
apply_anydoc_patch() {
    if grep -q 'asset://' "${ANYDOC_DIR}/src/render/markdown/inline.rs" 2>/dev/null; then
        echo "anydoc patch already applied"
        return
    fi
    if ! git -C "${ANYDOC_DIR}" apply --check "${ANYDOC_PATCH}" 2>/dev/null; then
        echo "error: anydoc patch no longer applies — regenerate ${ANYDOC_PATCH} (升级 anydoc tag 后常见)" >&2
        exit 2
    fi
    git -C "${ANYDOC_DIR}" apply "${ANYDOC_PATCH}"
    echo "applied ${ANYDOC_PATCH}"
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
    [ -f "${WASM}" ] || { echo "missing wasm — run: $0 wasm" >&2; exit 2; }
    # bin/ 不入库；缺失时自动从 third-party/goccy-wasm2go 构建（CI 也需要）。
    if [ ! -x "${ROOT}/bin/goccy-wasm2go" ]; then
        echo "building goccy-wasm2go → ${ROOT}/bin/goccy-wasm2go"
        (cd "${GOCCY_DIR}" && go build -o "${ROOT}/bin/goccy-wasm2go" ./cmd/wasm2go)
    fi
    # goccy AOT 生成多包 core（core.go + base/ + p0/ + p1/ + data.bin），
    # 整目录进 core/，不需要 split_gen 拆文件；生成耗时长（内部跑 go build
    # 抓 asm），但只发生在 wasm 变更后。
    rm -rf "${ROOT}/core"
    "${ROOT}/bin/goccy-wasm2go" -i "${WASM}" -pkg core -import anydoc-go/core \
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