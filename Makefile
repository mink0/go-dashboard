cwd := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
env_file := $(cwd)/.env
ifneq ("$(wildcard $(env_file))","")
	include $(env_file)
	export
endif

dev:
	air -c .air.toml

install:
	cd frontend && npm install
	go install

build-front:
	cd frontend && npm run build &&\
		statik -src=$(cwd)/frontend/dist -dest=$(cwd)/dist

build-go:
	go build -o $(cwd)/dist/server

build: build-front build-go

.PHONY: frontend
frontend serve:
	cd frontend && npm run serve

backend go-run:
	go run .

run: build-front backend
