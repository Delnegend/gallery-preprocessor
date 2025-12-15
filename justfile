@default:
	just --list

build:
	wails build

dev:
	wails dev

lint:
	#!/usr/bin/env bash
	go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -test ./...

	cd frontend && \
		bun x oxlint --import-plugin -D correctness -D perf --ignore-pattern wailsjs/**/*.* && \
		bun x prettier -l -w "**/*.{js,ts,vue,json,css}"