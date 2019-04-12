# tinygo

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/tinygo)](https://goreportcard.com/report/github.com/andygeiss/tinygo)

This project demonstrates how to shrink a WebAssembly App, with a total size of around 3 Megabyte, down to 50 Kilobyte by using tinygo. 

**Table of Contents**

- [Goals](README.md#goals)
- [Installation](README.md#installation)
- [Build](README.md#build)
- [Usage](README.md#usage)

## Goals

## Installation

**TinyGo via Docker**

    docker pull tinygo/tinygo:latest

**From Source**

    go get -u github.com/andygeiss/tinygo

## Build

**Build the WebAssembly file**

    sudo docker run \
        -v $PWD:/src \
        -v $GOPATH:/gohost \
        -e "GOPATH=$GOPATH:/gohost" \
        tinygo/tinygo \
        tinygo build -o /src/web/app/app.wasm -target wasm /src/main_wasm.go

**Compress the file**

    gzip -9 web/app/app.wasm \
        && mv web/app/app.wasm.gz web/app/app.wasm
         
## Usage

    sudo go run main.go

[http://127.0.0.1](http://127.0.0.1)
