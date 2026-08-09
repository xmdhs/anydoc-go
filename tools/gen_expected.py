#!/usr/bin/env python3
"""从 anydoc 的 insta 快照生成 testdata/expected 期望输出。

用法: gen_expected.py SNAPSHOTS_DIR EXPECTED_DIR [CASE...]
CASE 形如 `docx/text`（对应 fixtures/docx/text.docx 与
snapshots__docx__text.docx.snap）；缺省处理内置列表。
"""
import os
import re
import sys

SRC = sys.argv[1]
DST = sys.argv[2]

CASES = [
    "docx/text.docx",
    "csv/sheet.csv",
    "doc/text.doc",
    "pptx/pres.pptx",
    "epub/book.epub",
    "ods/sheet.ods",
]


def snap_name(case: str) -> str:
    path, ext = os.path.splitext(case)
    path = path.replace("/", "__")
    return f"snapshots__{path}{ext}.snap"


def main() -> None:
    cases = sys.argv[3:] or CASES
    for case in cases:
        snap = open(os.path.join(SRC, snap_name(case)), encoding="utf-8").read()
        body = snap.split("---\n")[-1].rstrip("\n")
        out_path = os.path.join(DST, os.path.splitext(case)[0] + ".md")
        os.makedirs(os.path.dirname(out_path), exist_ok=True)
        with open(out_path, "w", encoding="utf-8") as f:
            f.write(body + ("\n" if body and not body.endswith("\n") else ""))
        print(f"{case} -> {out_path}", file=sys.stderr)


if __name__ == "__main__":
    main()