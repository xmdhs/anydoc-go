#!/usr/bin/env python3
"""用 reportlab + Pillow 生成带图 PDF 夹具。

全部缓存/临时文件落项目内：UV_CACHE_DIR=.uv-cache, TMPDIR=.tmp
执行：UV_CACHE_DIR=$PWD/.uv-cache TMPDIR=$PWD/.tmp uv run --with reportlab --with Pillow python tools/gen_pdf_fixtures.py

产物：
  testdata/fixtures/pdf/with-text-image.pdf  文字 + 1 张 JPEG (DCTDecode)
  testdata/fixtures/pdf/scanned-image.pdf    无文字仅 1 张 JPEG (扫描件)
  testdata/fixtures/pdf/multi-image.pdf      文字 + 2 图 (JPEG + PNG/Flate)
"""
import pathlib

ROOT = pathlib.Path(__file__).resolve().parent.parent
OUT_DIR = ROOT / "testdata" / "fixtures" / "pdf"
TMP_DIR = ROOT / ".tmp"

OUT_DIR.mkdir(parents=True, exist_ok=True)
TMP_DIR.mkdir(parents=True, exist_ok=True)

from PIL import Image as PILImage
from reportlab.lib.pagesizes import A4
from reportlab.lib.utils import ImageReader
from reportlab.pdfgen import canvas

W, H = A4


def make_jpeg(path: pathlib.Path, color=(220, 40, 40)):
    img = PILImage.new("RGB", (128, 128), color)
    # draw a simple pattern so it is not solid
    for x in range(0, 128, 16):
        for y in range(0, 128, 16):
            if (x + y) % 32 == 0:
                for dx in range(8):
                    for dy in range(8):
                        if x + dx < 128 and y + dy < 128:
                            img.putpixel((x + dx, y + dy), (255, 255, 255))
    img.save(path, "JPEG", quality=85)
    return path


def make_png(path: pathlib.Path, color=(40, 80, 200)):
    img = PILImage.new("RGB", (128, 128), color)
    for x in range(0, 128, 16):
        for y in range(0, 128, 16):
            if (x + y) % 32 == 0:
                for dx in range(8):
                    for dy in range(8):
                        if x + dx < 128 and y + dy < 128:
                            img.putpixel((x + dx, y + dy), (255, 220, 0))
    img.save(path, "PNG")
    return path


jpeg_path = TMP_DIR / "_gen_red.jpg"
png_path = TMP_DIR / "_gen_blue.png"
make_jpeg(jpeg_path, (220, 40, 40))
make_png(png_path, (40, 80, 200))

# 1. with-text-image.pdf — 文字 + 1 JPEG
out1 = OUT_DIR / "with-text-image.pdf"
c = canvas.Canvas(str(out1), pagesize=A4)
c.setFont("Helvetica", 12)
c.drawString(50, 750, "Fixture Document with Image")
c.drawString(50, 730, "This PDF contains extractable text and one embedded JPEG image.")
c.drawImage(ImageReader(str(jpeg_path)), 50, 500, width=200, height=200, preserveAspectRatio=True, mask="auto")
c.setFont("Helvetica", 9)
c.drawString(50, 480, "Image above is a red square with white pattern (JPEG DCTDecode).")
c.showPage()
c.save()
print(f"wrote {out1} ({out1.stat().st_size} bytes)")

# 2. scanned-image.pdf — 无文字仅 1 JPEG (扫描件)
out2 = OUT_DIR / "scanned-image.pdf"
c = canvas.Canvas(str(out2), pagesize=A4)
c.drawImage(ImageReader(str(jpeg_path)), 50, 400, width=400, height=400, preserveAspectRatio=True, mask="auto")
c.showPage()
c.save()
print(f"wrote {out2} ({out2.stat().st_size} bytes)")

# 3. multi-image.pdf — 文字 + 2 图 (JPEG + PNG)
out3 = OUT_DIR / "multi-image.pdf"
c = canvas.Canvas(str(out3), pagesize=A4)
c.setFont("Helvetica", 12)
c.drawString(50, 750, "Multi Image Fixture")
c.drawString(50, 730, "This PDF has two images: JPEG and PNG (FlateDecode).")
c.drawImage(ImageReader(str(jpeg_path)), 50, 500, width=180, height=180, preserveAspectRatio=True, mask="auto")
c.drawImage(ImageReader(str(png_path)), 300, 500, width=180, height=180, preserveAspectRatio=True, mask="auto")
c.setFont("Helvetica", 9)
c.drawString(50, 480, "Left: JPEG DCTDecode. Right: PNG FlateDecode.")
c.showPage()
c.save()
print(f"wrote {out3} ({out3.stat().st_size} bytes)")

# quick verify: check filters in output
for p in [out1, out2, out3]:
    data = p.read_bytes()
    filters = [x.decode() for x in [b"DCTDecode", b"FlateDecode", b"JPXDecode"] if x in data]
    cnt = data.count(b"/Subtype /Image") + data.count(b"/Subtype/Image")
    print(f"  {p.name}: images~{cnt} filters={filters} size={len(data)} Tj={'Tj' in data.decode(errors='ignore')}")
