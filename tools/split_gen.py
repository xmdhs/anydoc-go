#!/usr/bin/env python3
"""Split the wasm2go-generated anydoc.wasm.go into smaller per-chunk files
so `go build` fits in limited memory (one giant file exceeds the toolchain's
memory budget).

Usage: split_gen.py [FUNCS_PER_FILE] [WORKDIR]
Reads WORKDIR/anydoc.wasm.go, writes WORKDIR/anydoc_gen_NNN.go, removes the
original.

The head file keeps package/imports/types/consts/go:embed var; chunks get
only the imports their bodies actually use.
"""

import os
import re
import sys

FUNCS_PER_FILE = int(sys.argv[1]) if len(sys.argv) > 1 else 150
WORKDIR = sys.argv[2] if len(sys.argv) > 2 else os.path.dirname(os.path.abspath(__file__))
SRC = os.path.join(WORKDIR, "anydoc.wasm.go")

lines = open(SRC, encoding="utf-8").read().split("\n")

# Head: everything before the first top-level func.
first = next(i for i, l in enumerate(lines) if re.match(r"^func(?:\s|\(m \*Module\))", l))
head = lines[:first]
rest = lines[first:]

# Split the rest into top-level declaration elements. An element is opened by
# any line (blank or //go: directive) and closed only once a `func` header has
# been seen and its braces balance, so directives stay glued to their
# declaration.
funcs = []
cur = []
seen = False
depth = 0
for l in rest:
    if not cur:
        cur = [l]
        seen = l.startswith("func")
        if seen:
            depth = l.count("{") - l.count("}")
        continue
    cur.append(l)
    if l.startswith("func") and not seen:
        seen = True
        depth = l.count("{") - l.count("}")
    elif seen:
        depth += l.count("{") - l.count("}")
    if seen and depth == 0:
        funcs.append("\n".join(cur))
        cur, seen, depth = [], False, 0
if cur:
    funcs.append("\n".join(cur))

head_text = "\n".join(head) + "\n"
head_imports = re.findall(r'^\t(?:_ )?"([^"]+)"$', head_text, re.M)

os.remove(SRC)
# The head file has no function bodies; keep only the go:embed import.
head_text = re.sub(r'import \(\n(?:\t[^\n]*\n)+\)\n', 'import (\n\t_ "embed"\n)\n', head_text, count=1)
m = re.search(r"^package (\w+)", head_text, re.M)
PKG = m.group(1) if m else "main"
with open(os.path.join(WORKDIR, "anydoc_gen_000.go"), "w", encoding="utf-8") as f:
    f.write(head_text)

def needed_imports(body):
    used = []
    if "//go:embed" in body:
        used.append('_ "embed"')
    if "binary." in body:
        used.append('"encoding/binary"')
    if "bits." in body:
        used.append('"math/bits"')
    if "math." in body:
        used.append('"math"')
    if "unsafe." in body:
        used.append('"unsafe"')
    if "runtime." in body:
        used.append('"runtime"')
    return sorted(used)

chunks = [funcs[i : i + FUNCS_PER_FILE] for i in range(0, len(funcs), FUNCS_PER_FILE)]
for idx, chunk in enumerate(chunks, start=1):
    body = "\n".join(chunk)
    used = needed_imports(body)
    with open(os.path.join(WORKDIR, f"anydoc_gen_{idx:03d}.go"), "w", encoding="utf-8") as f:
        f.write(f"package {PKG}\n\n")
        if used:
            f.write("import (\n")
            for p in used:
                f.write(f"\t{p}\n")
            f.write(")\n\n")
        f.write(body + "\n")

print(f"split {len(funcs)} top-level funcs into {1 + len(chunks)} files", file=sys.stderr)