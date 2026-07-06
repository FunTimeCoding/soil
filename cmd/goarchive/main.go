package main

import "github.com/funtimecoding/soil/pkg/tool/goarchive"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goarchive.Main(Version, GitHash, BuildDate)
}
