@default:
	just --choose

build:
	wails build

dev:
	wails dev

check:
	#!/usr/bin/env bash

	go fmt
	go vet

	cd frontend && \
		bun x oxlint --import-plugin -D correctness -D perf --ignore-pattern wailsjs/**/*.* && \
		bun x prettier -l -w "**/*.{js,ts,vue,json,css}"

# Manifest-less: version lives in git tags (frontend/package.json stays 0.0.0)
bump version:
	@echo "versions are tracked by git tags; nothing to bump"
