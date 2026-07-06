package main

import "github.com/funtimecoding/soil/pkg/tool/gochk"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gochk.Main(Version, GitHash, BuildDate)
}
