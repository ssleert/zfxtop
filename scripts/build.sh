#!/bin/sh
. './scripts/lib.sh'

CGO_ENABLED=0
GOFLAGS="-ldflags='-s -w' -gcflags=all='-B -C'"

main() {
  generate
  go build cmd/zfxtop/zfxtop.go
}

main
