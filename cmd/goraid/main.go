package main

import "github.com/funtimecoding/soil/pkg/tool/goraid"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goraid.Main(Version, GitHash, BuildDate)
}
