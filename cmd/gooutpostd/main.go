package main

import "github.com/funtimecoding/soil/pkg/tool/gooutpostd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gooutpostd.Main(Version, GitHash, BuildDate)
}
