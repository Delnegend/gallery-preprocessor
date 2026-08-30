#!/usr/bin/env bash
set -e
# Toolchain (go/wails/bun/just) is pre-installed in Dockerfile for cache + Zed.
# This hook only handles workspace-mounted deps that cannot be baked into the image
# (bind mount at /workspaces/gallery-preprocessor hides image's frontend/node_modules).

# Ensure wails matches go.mod if image is stale (no rebuild after `go.mod` bump)
WAILS_VERSION="$(grep -E '^\s*github.com/wailsapp/wails/v2\s+v' go.mod | awk '{print $2}')"
if [ -n "$WAILS_VERSION" ] && ! wails version 2>/dev/null | grep -q "$WAILS_VERSION"; then
  echo "wails $WAILS_VERSION not found, installing..."
  go install "github.com/wailsapp/wails/v2/cmd/wails@$WAILS_VERSION"
fi

# frontend deps (must run on mounted workspace)
cd frontend && bun i
