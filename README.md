# go-dashboard

Cross-platform single-page application prototype which could be compiled into a single binary using [Golang](https://go.dev/).
The application consists of two parts: the front-end SPA with realtime data display written in VueJS and the back-end API web server written in Go.
This approach allows you to minify and pack any complex assents into a single executable file which could.

## Requirements

- [go](https://go.dev/dl/)
- [nodejs](https://nodejs.org/en/)
- [statik](https://github.com/rakyll/statik)

### Optional

- [air](https://github.com/cosmtrek/air) for live reload

## Install

    make install

## Run

    make run

## Build

    make build

## Dev environment with live-reload

    make dev
