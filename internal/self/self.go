package self

import _ "embed"

//go:generate sh -c "git rev-parse --short HEAD > commit.dat"
//go:embed commit.dat
var Commit string

//go:generate sh -c "cat ../../VERSION > version.dat"
//go:embed version.dat
var Version string
