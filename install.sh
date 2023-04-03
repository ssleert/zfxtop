#!/bin/sh

export CLR_RESET='\033[1;0m'
export STL_BOLD='\033[1;1m'
export CLR_RED='\033[0;31m'
export CLR_GREEN='\033[0;32m'
export CLR_BLUE='\033[0;34m'

has() {
  _cmd=$(command -v "$1") 2> /dev/null || return 1
  [ -x "$_cmd" ] || return 1
}

msg() {
  printf "${CLR_BLUE}${STL_BOLD}::${CLR_RESET} ${STL_BOLD}%s${CLR_RESET}\n" "$1"
  shift
  for i in "$@"; do
    printf " ${CLR_BLUE}${STL_BOLD}|${CLR_RESET} ${STL_BOLD}%s${CLR_RESET}\n" "$i"
  done
}

compl() {
  printf "${CLR_GREEN}${STL_BOLD}>>>${CLR_RESET} ${STL_BOLD}%s${CLR_RESET}\n" "$1"
  shift
  for i in "$@"; do
    printf "  ${CLR_GREEN}${STL_BOLD}|${CLR_RESET} ${STL_BOLD}%s${CLR_RESET}\n" "$i"
  done
}

die() {
  printf "${CLR_RED}${STL_BOLD}error:${CLR_RESET} ${STL_BOLD}%s${CLR_RESET}\n" "$1"
  shift
  for i in "$@"; do
    printf "     ${CLR_RED}${STL_BOLD}|${CLR_RESET} ${STL_BOLD}%s${CLR_RESET}\n" "$i"
  done
  exit 1
}

download() {
  curl -sfLo "$1" "$2"
}

verify_checksums() {
  msg 'Verifying checksums.'
  if has sha256sum; then
    _ok=$(sha256sum --ignore-missing --quiet --check checksums.txt)
  else
    _ok=$(shasum -a 256 --ignore-missing --quiet --check checksums.txt)
  fi

  $_ok || die 'Checksums did not match! Abort'
}

main() {
  NAME="zfxtop"

  OS="$(uname -s)"
  msg 'OS fetched.'

  ARCH="$(uname -m)"
  msg 'ARCH fetched.'

  TMPDIR="$(mktemp -d)"
  msg 'Temp dir created.'

  compl 'All required info fetched.'

  RELEASES_URL="https://github.com/ssleert/${NAME}/releases"
  msg 'Fetching latest version'
  TAG="$(curl -sfL -o /dev/null -w "%{url_effective}" "${RELEASES_URL}/latest" | \
    rev | \
    cut -f1 -d'/' | \
    rev)"

  [ -z "$TAG" ] && {
    die "Unable to get ${NAME} version." \
        'try to install curl'
  }

  echo "$TAG" | grep -qE '^v[0-9]+\.[0-9]+\.[0-9]+$' && {
    die "Unable to get ${NAME} version."
  }

  VERSION="$TAG"
  compl "Latest version is ${VERSION}"

  case "$ARCH" in
    "x86_64")
      ARCH="amd64"
    ;;
    i*)
      ARCH="386"
    ;;
    *)
      die "$ARCH is unsupported"
    ;;
  esac
  msg 'ARCH configured.'

  case "$OS" in
    "Linux")
      OS="linux"
    ;;
    *)
      die "$OS is unsupported"
    ;;
  esac
  msg 'OS configured.'

  export TAR_NAME="${NAME}_${VERSION}_${OS}_${ARCH}.tar.gz"
  export TAR_FILE="${TMPDIR}/${TAR_NAME}"

  (
    cd "$TMPDIR" || return

    msg "Downloading $TAR_NAME"
    download "$TAR_FILE" "$RELEASES_URL/download/$TAG/$TAR_NAME"

    msg "Downloading checksums"
    download "checksums.txt" "$RELEASES_URL/download/$TAG/checksums.txt"

    verify_checksums
  )

  msg "Extract ${TAR_FILE}"
  tar -xzf "$TAR_FILE" -C "$TMPDIR/"

  if has zfxtop; then
    OUT="$(dirname "$(command -v zfxtop)")"
  else
    OUT="/usr/local/bin"
  fi

  mkdir -p "$OUT"
  msg "Moving to '/usr/bin'"

  if has doas; then
    doas mv "${TMPDIR}/${NAME}_${VERSION}_${OS}_${ARCH}" "${OUT}/${NAME}"
  elif has sudo; then
    sudo mv "${TMPDIR}/${NAME}_${VERSION}_${OS}_${ARCH}" "${OUT}/${NAME}"
  fi

  compl "Installation COMPLETED."
}

main
