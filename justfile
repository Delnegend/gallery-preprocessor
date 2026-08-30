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
