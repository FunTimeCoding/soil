package main

import "github.com/funtimecoding/soil/pkg/tool/gowiki"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gowiki.Main(Version, GitHash, BuildDate)
}
