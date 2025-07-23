@default:
	just --list

build:
	wails build -upx

dev:
	wails dev

lint:
	#!/usr/bin/env bash
	go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -test ./...

	cd frontend && \
		pnpm oxlint --import-plugin -D correctness -D perf \
		--ignore-pattern wailsjs/**/*.* && \
		pnpm prettier -l -w "**/*.{js,ts,vue,json,css}"