build:
    wails build -upx

dev:
    wails dev

lint:
    go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -test ./...