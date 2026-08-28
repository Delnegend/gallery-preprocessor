#!/usr/bin/env bash

# just
echo 'eval "$(just --completions bash)"' >> ~/.bashrc

# go
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# bun
curl -fsSL https://bun.sh/install | bash
echo 'export BUN_INSTALL="$HOME/.bun"' >> ~/.bashrc
echo 'export PATH="$BUN_INSTALL/bin:$PATH"' >> ~/.bashrc
