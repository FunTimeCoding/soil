package main

import "github.com/funtimecoding/soil/pkg/tool/goalpine"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goalpine.Main(Version, GitHash, BuildDate)
}
