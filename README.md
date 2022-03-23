# go-dashboard

Cross-platform single-page application prototype which could be compiled into a single binary using [Golang](https://go.dev/).
The application consists of two parts: the front-end SPA with realtime data display written in VueJS and the back-end API web server written in Go.
This approach allows you to minify and pack any complex assents into a single executable file which could.

## Requirements

- [go](https://go.dev/dl/)
- [nodejs](https://nodejs.org/en/)

### Optional

- [air](https://github.com/cosmtrek/air) for live reload

## Install

Install the command line tool for [statik](https://github.com/rakyll/statik) first:
    
    go install github.com/rakyll/statik@v0.1.7

Statik is a tiny program that reads a directory and generates a source file that contains its contents. The generated source file registers the directory contents to be used by statik file system.

Then run make install:

    make install

## Run

    make run

## Build

    make build

## Dev environment with live-reload

    make dev
