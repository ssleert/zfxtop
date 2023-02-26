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

err() {
  printf "${CLR_RED}${STL_BOLD}error:${CLR_RESET} ${STL_BOLD}%s${CLR_RESET}\n" "$1"
  shift
  for i in "$@"; do
    printf "     ${CLR_RED}${STL_BOLD}|${CLR_RESET} ${STL_BOLD}%s${CLR_RESET}\n" "$i"
  done
}

die() {
  err "$@"
  exit 1
}

generate() {
  go generate ./...
  compl "files generated"
}

format() {
  go fmt ./...
  compl "files formated"
}
