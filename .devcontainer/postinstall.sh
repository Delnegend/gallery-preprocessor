#!/usr/bin/env bash

# just
echo 'eval "$(just --completions bash)"' >> ~/.bashrc

# go - install wails CLI matching go.mod to avoid version mismatch warning
WAILS_VERSION="$(grep -E '^\s*github.com/wailsapp/wails/v2\s+v' go.mod | awk '{print $2}')"
if [ -n "$WAILS_VERSION" ]; then
  go install "github.com/wailsapp/wails/v2/cmd/wails@$WAILS_VERSION"
else
  go install github.com/wailsapp/wails/v2/cmd/wails@latest
fi

# bun
curl -fsSL https://bun.sh/install | bash
echo 'export BUN_INSTALL="$HOME/.bun"' >> ~/.bashrc
echo 'export PATH="$BUN_INSTALL/bin:$PATH"' >> ~/.bashrc
# make bun available to non-interactive shells (wails executes `bun` via $PATH)
sudo ln -sf "$HOME/.bun/bin/bun" /usr/local/bin/bun
