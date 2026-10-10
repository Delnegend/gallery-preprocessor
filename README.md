<div align="center">

# gallery-preprocessor

**Drag-and-drop desktop app that batch-processes comic and manga pages: artifact removal, JPEG-XL and AVIF conversion, 7z parity files, and reversible diff sequences.**

[![CI](https://img.shields.io/github/actions/workflow/status/Delnegend/gallery-preprocessor/ci.yml?branch=main&style=flat-square)](https://github.com/Delnegend/gallery-preprocessor/actions)
[![Release](https://img.shields.io/github/v/release/Delnegend/gallery-preprocessor?style=flat-square)](https://github.com/Delnegend/gallery-preprocessor/releases)
[![License](https://img.shields.io/github/license/Delnegend/gallery-preprocessor?style=flat-square)](LICENSE)

</div>

---

> **Archived.** Development stopped in October 2026 and this repository is read-only. The
> published releases stay available, and the source still builds as described below.

## Quick Start

```bash
# 1. Download (Linux x64 — macOS and Windows archives are on the Releases page)
curl -LO https://github.com/Delnegend/gallery-preprocessor/releases/latest/download/gallery-preprocessor-linux-amd64.tar.xz
# 2. Extract
tar -xJf gallery-preprocessor-linux-amd64.tar.xz
# 3. Run it
./gallery-preprocessor-linux-amd64/gallery-preprocessor
```

Drop files or a whole folder onto one of the task tiles to start.

## Highlights

- **Seven pipelines, no configuration** — every task declares the file types it accepts and streams its progress into one log pane.
- **Reversible by design** — `Differ diff` turns a folder into small change-only PNGs, and `Differ join` rebuilds the images from them.
- **Safe to interrupt** — long batches report progress and warnings continuously, and can be cancelled mid-run.
- **No runtime to install** — a single desktop binary for Linux, macOS and Windows.

## Tasks

| Tile | What it does | Accepts |
|---|---|---|
| Artefact | Remove JPEG compression artifacts, output PNG | `.jpg` |
| Artefact + AVIF (Lossy) | Artifact removal, then AVIF lossy compression | `.jpg` |
| CJXL (Lossless) | Compress JPG/PNG to JPEG-XL, lossless | `.jpg`, `.png` |
| AVIF (Lossy) | Compress JPG/PNG to AVIF | `.jpg`, `.png` |
| DJXL | Decompress JPEG-XL back to JPG/PNG | `.jxl` |
| PAR2 | Create parity files for a 7z archive | `.7z` |
| Differ diff / join | Generate, then reassemble, PNG diff sequences | `.png` |

## Build From Source

```bash
git clone https://github.com/Delnegend/gallery-preprocessor && cd gallery-preprocessor
just build   # wails build -tags webkit2_41 on Linux
just check   # go fmt + go vet + oxlint + prettier
```

Needs Go (see `go.mod`), [Bun](https://bun.sh), the [Wails CLI](https://wails.io/docs/gettingstarted/installation) and, on Linux, `libgtk-3-dev libwebkit2gtk-4.1-dev pkg-config build-essential`. A devcontainer is included — *Reopen in Container* installs everything.

## License

[MIT](LICENSE)
