#!/bin/sh
. './scripts/lib.sh'

CGO_ENABLED=0
GOFLAGS="-ldflags='-s -w' -gcflags=all='-B -C'"

NAME='zfxtop'
VERSION="$(cat VERSION)"
RELEASE='./release'
BDIR="${RELEASE}/bin"
PDIR="${RELEASE}/pkg"

ARCHES='
  386
  amd64
'
OSES='
  linux
'

clear() {
  [ -d "$RELEASE" ] && {
    rm -r "$RELEASE"
    msg "'$RELEASE' dir removed"
  }

  mkdir "$RELEASE"
  msg "'$RELEASE' dir created"
}

compile() {
  mkdir "$BDIR"
  msg "'$BDIR' created"

  for arch in $ARCHES; do
    for os in $OSES; do
      GOARCH="$arch" GOOS="$os" go build -o "${BDIR}/${NAME}_${VERSION}_${os}_${arch}" "cmd/${NAME}/${NAME}.go"
      [ $? -ne 0 ] && {
        die 'exeption during compilation'
      }
      msg "'${BDIR}/${NAME}_${VERSION}_${os}_${arch}' builded"
    done
  done
  compl 'all executables builded'
}

package() {
  mkdir "$PDIR"
  msg "'$PDIR' created"

  for arch in $ARCHES; do
    for os in $OSES; do
      tar -C "${BDIR}/" -czf "${PDIR}/${NAME}_${VERSION}_${os}_${arch}.tar.gz" "${NAME}_${VERSION}_${os}_${arch}"
      [ $? -ne 0 ] && {
        die 'exeption during packaging'
      }
      msg "'${PDIR}/${NAME}_${VERSION}_${os}_${arch}.tar.gz' packaged"
    done
  done
  compl 'all executables packaged'
}

gen_checksum() {
  cd "$PDIR"
  msg "dir changed to '$PDIR'"
  sha256sum *.tar.gz > checksums.txt
  compl "checksums for '$PDIR' generated"
  cd "../../"
  msg "dir chenged to '$PWD'"
}

main() {
  clear
  generate
  compile
  package
  gen_checksum
  compl 'ALL COMPLETED'
}

main
