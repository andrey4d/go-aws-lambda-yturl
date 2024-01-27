#!/bin/sh

GOOS="linux"
GOARCH="amd64"

BIN="bin/hello"


if [ $GOOS = "windows" ]; then
	output_name+=".exe"
fi

env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o "${BIN}" cmd/main.go  
zip ${BIN}.zip ${BIN}
