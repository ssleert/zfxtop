CGO_ENABLED=0

main() {
  generate
  go build -ldflags="-s -w" -gcflags=all="-B -C" cmd/zfxtop/zfxtop.go
}

main
