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
export RUSTUP_HOME="${RUSTUP_HOME:-${ROOT}/../.rustup}"
export CARGO_HOME="${CARGO_HOME:-${ROOT}/../.cargo}"
export CARGO_TARGET_DIR="${CARGO_TARGET_DIR:-${ROOT}/target}"
export ZIG_GLOBAL_CACHE_DIR="${ZIG_GLOBAL_CACHE_DIR:-${ROOT}/.zig-cache}"
export ZIG_LOCAL_CACHE_DIR="${ZIG_LOCAL_CACHE_DIR:-${ROOT}/libzig-cache}"
export GOCACHE="${GOCACHE:-${ROOT}/../.go-cache}"
export GOPATH="${GOPATH:-${ROOT}/../.gopath}"
export PATH="${CARGO_HOME}/bin:${ROOT}/bin:${PATH}"

ANYDOC_DIR="${ROOT}/third-party/anydoc"
WASM2GO_DIR="${ROOT}/third-party/wasm2go"
ANYDOC_REPO="${ANYDOC_REPO:-https://github.com/firecrawl/anydoc}"
WASM2GO_REPO="${WASM2GO_REPO:-https://github.com/ncruces/wasm2go}"

FETCH_REF="${FETCH_REF:-main}"      # fetch 时拉的分支/标签
WASM="${CARGO_TARGET_DIR}/wasm32-unknown-unknown/release/anydoc_cabi.wasm"

fetch() {
    # 第三方仓库位于 anydoc-go 自己的 third-party/ 下（.gitignore 排除，
    # 不随本仓库提交）。cabi 对 anydoc 有本地改动（pdf feature 开关），
    # fetch 只做 fetch/checkout/ff-only pull，有本地改动且冲突时不强推。
    fetch_one anydoc   "${ANYDOC_REPO}"   "${ANYDOC_DIR}"
    fetch_one wasm2go  "${WASM2GO_REPO}"  "${WASM2GO_DIR}"
    echo "third-party repos ready:" "${ANYDOC_DIR}" "${WASM2GO_DIR}"
}

fetch_one() {
    local name="$1" url="$2" dir="$3"
    if [ ! -d "${dir}/.git" ]; then
        echo "cloning ${name} (${url}) → ${dir}"
        git clone -q -b "${FETCH_REF}" "${url}" "${dir}"
    else
        echo "updating ${name} → ${dir}"
        git -C "${dir}" fetch -q origin || true
        git -C "${dir}" checkout -q "${FETCH_REF}" 2>/dev/null || true
        git -C "${dir}" pull -q --ff-only origin "${FETCH_REF}" || true
    fi
}

require_repos() {
    [ -d "${ANYDOC_DIR}/.git" ] || { echo "missing third-party/anydoc — run: $0 fetch" >&2; exit 2; }
    [ -d "${WASM2GO_DIR}/.git" ] || { echo "missing third-party/wasm2go — run: $0 fetch" >&2; exit 2; }
}

case "${1:-build}" in
    fetch) fetch ;;
    wasm)
        require_repos
        cd "${ROOT}/cabi"
        cargo build --release --target wasm32-unknown-unknown
        ;;
    cli)
        [ -f "${WASM}" ] || { echo "missing wasm — run: $0 wasm" >&2; exit 2; }
        [ -x "${ROOT}/bin/wasm2go" ] || { echo "missing bin/wasm2go — run: $0 fetch && (cd third-party/wasm2go && go build -o ${ROOT}/bin/wasm2go .)" >&2; exit 2; }
        mkdir -p "${ROOT}/core"
        "${ROOT}/bin/wasm2go" -pkg core -embed -o "${ROOT}/core/anydoc.wasm.go" "${WASM}"
        python3 "${ROOT}/tools/split_gen.py" 45 "${ROOT}/core"
        cd "${ROOT}"
        go build -p 2 -gcflags=all="-N -l" -o "${ROOT}/bin/anydoc" .
        ;;
    test)
        cd "${ROOT}"
        go test -p 1 -gcflags=all="-N -l" -v ./...
        ;;
    *)
        echo "usage: $0 [fetch|wasm|cli|test|build]" >&2
        exit 2
        ;;
esac