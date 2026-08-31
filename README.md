# gallery-preprocessor

![Go](https://img.shields.io/badge/Go-1.27-00ADD8?style=flat&logo=go&logoColor=white)
![Wails](https://img.shields.io/badge/Wails-v2.15.0-EF4444?style=flat)
![Vue](https://img.shields.io/badge/Vue-3.5-42B883?style=flat&logo=vue.js&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-8.2-646CFF?style=flat&logo=vite&logoColor=white)
![Bun](https://img.shields.io/badge/Bun-1.x-000000?style=flat&logo=bun&logoColor=white)
![TailwindCSS](https://img.shields.io/badge/Tailwind-4.3-06B6D4?style=flat&logo=tailwindcss&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-7.0-3178C6?style=flat&logo=typescript&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-yellow?style=flat)
![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey?style=flat)

Drag-and-drop desktop app to batch-process gallery images — artifact removal, JPEG-XL and AVIF conversion, and diff/parity helpers. Built with [Wails](https://wails.io) (Go backend + Vue frontend).

## Features

- **Artefact** — remove JPEG compression artifacts, output PNG. Accepts `.jpg`
- **Artefact + AVIF (Lossy)** — artefact removal + AVIF compression. Accepts `.jpg`
- **CJXL (Lossless)** — compress JPG/PNG to JPEG-XL lossless. Accepts `.jpg`, `.png`
- **AVIF (Lossy)** — compress JPG/PNG to AVIF lossy. Accepts `.jpg`, `.png`
- **DJXL** — decompress JPEG-XL to JPG/PNG. Accepts `.jxl`
- **PAR2** — create parity files for `.7z` archives. Accepts `.7z`
- **Differ diff / join** — generate and re-assemble PNG diff sequences. Accepts `.png`

All tasks show live progress and warnings in the log pane; long-running tasks can be cancelled.

## Quick start

### Installation

Download the latest pre-built binary from [Releases](https://github.com/Delnegend/gallery-preprocessor/releases) (output after local build is `build/bin/gallery-preprocessor`, `.exe` on Windows).

### Usage

1. Launch the app (double-click the binary or `wails dev`).
2. Drag files or a folder onto one of the 8 task tiles.
3. Monitor `Progress` and `Warning` output; use **Cancel Task** to abort.

Tasks and accepted extensions are defined in `frontend/src/App.vue:25` and `backend/PerformTask.go:15`.

## Development

### Prerequisites

**Recommended: devcontainer** — no host toolchain install needed:

```bash
# VS Code: Command Palette → Reopen in Container
# CLI:
devcontainer up --workspace-folder .
```

**Without devcontainer:**

- [Go](https://go.dev) (see `go.mod`), [Bun](https://bun.sh), [Wails](https://wails.io/docs/gettingstarted/installation)
- Linux: `libgtk-3-dev` `libwebkit2gtk-4.1-dev` `pkg-config` `build-essential`

### Build

```bash
git clone https://github.com/Delnegend/gallery-preprocessor.git
cd gallery-preprocessor

# dev with hot reload
wails dev  # or: just dev

# production build
wails build -upx
# Linux requires webkit tag (already set in wails.json:5)
wails build -tags webkit2_41 -upx  # or: just build

# frontend only
cd frontend && bun i && bun run dev
cd frontend && bun run build
```

### Checks

```bash
just check   # go fmt + go vet + oxlint + prettier
```

## License

MIT
fake dependabot test
