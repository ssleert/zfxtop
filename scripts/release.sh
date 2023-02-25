CGO_ENABLED=0

go generate ./...
GOARCH=386 go build -ldflags="-s -w" -gcflags=all="-B -C" -o "release/zfxtop_386" cmd/zfxtop/zfxtop.go
GOARCH=amd64 go build -ldflags="-s -w" -gcflags=all="-B -C" -o "release/zfxtop_amd64" cmd/zfxtop/zfxtop.go
# sorry but now without arm support(
