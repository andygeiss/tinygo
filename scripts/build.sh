#!/bin/bash

sudo docker run \
    -v $PWD:/src \
    -v $GOPATH:/gohost \
    -e "GOPATH=$GOPATH:/gohost" \
    tinygo/tinygo \
    tinygo build -o /src/web/app/app.wasm -target wasm /src/main_wasm.go

gzip -9 web/app/app.wasm
mv web/app/app.wasm.gz web/app/app.wasm
