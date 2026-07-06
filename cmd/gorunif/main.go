package main

import "github.com/funtimecoding/soil/pkg/tool/gorunif"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gorunif.Main(Version, GitHash, BuildDate)
}
