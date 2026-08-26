package main

import "github.com/funtimecoding/soil/pkg/tool/gohw"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gohw.Main(Version, GitHash, BuildDate)
}
