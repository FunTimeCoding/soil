package main

import "github.com/funtimecoding/soil/pkg/tool/goflightd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goflightd.Main(Version, GitHash, BuildDate)
}
