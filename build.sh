#!/usr/bin/env bash
# anydoc-go 构建管线：
#   fetch: 拉取/刷新第三方仓库（third-party/anydoc、third-party/wasm2go），
#          URL 可用 ANPYOC_REPO / WASM2GO_REPO 覆盖
#   wasm:  cargo → target/<arch>/release/anydoc_cabi.wasm（依赖 third-party/anydoc）
#   cli:   wasm2go → 拆分 core/ → go build → bin/anydoc
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
WASM2GO_DIR="${ROOT}/third-party/wasm2go"
ANYDOC_REPO="${ANYDOC_REPO:-https://github.com/firecrawl/anydoc}"
WASM2GO_REPO="${WASM2GO_REPO:-https://github.com/ncruces/wasm2go}"

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
    # 也便于我们以 tag 为准复现构建）；wasm2go 跟 main。
    # fetch 只做 fetch/checkout/ff-only pull，有本地改动且冲突时不强推。
    fetch_one anydoc   "${ANYDOC_REPO}"   "${ANYDOC_DIR}"
    fetch_one wasm2go  "${WASM2GO_REPO}"  "${WASM2GO_DIR}"
    echo "third-party repos ready:" "${ANYDOC_DIR}" "${WASM2GO_DIR}"
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
    [ -d "${WASM2GO_DIR}/.git" ] || { echo "missing third-party/wasm2go — run: $0 fetch" >&2; exit 2; }
}

wasm_step() {
    require_repos
    cd "${ROOT}/cabi"
    cargo build --release --target wasm32-unknown-unknown
    cd "${ROOT}"
}

cli_step() {
    [ -f "${WASM}" ] || { echo "missing wasm — run: $0 wasm" >&2; exit 2; }
    [ -x "${ROOT}/bin/wasm2go" ] || { echo "missing bin/wasm2go — run: $0 fetch && (cd third-party/wasm2go && go build -o ${ROOT}/bin/wasm2go .)" >&2; exit 2; }
    mkdir -p "${ROOT}/core"
    rm -f "${ROOT}/core/anydoc_gen_*.go" "${ROOT}/core/anydoc.wasm.go"
    # -unsafe 用指针直读替代 encoding/binary 解码，转换快约 1.1~1.4 倍
    # （基准见 bench_test.go）；生成代码仍保持边界检查语义。
    "${ROOT}/bin/wasm2go" -unsafe -pkg core -embed -o "${ROOT}/core/anydoc.wasm.go" "${WASM}"
    python3 "${ROOT}/tools/split_gen.py" 45 "${ROOT}/core"
    cd "${ROOT}"
    go build -p 2 -gcflags=all="-N -l" -o "${ROOT}/bin/anydoc" .
}

test_step() {
    cd "${ROOT}"
    go test -p 1 -gcflags=all="-N -l" -v ./...
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