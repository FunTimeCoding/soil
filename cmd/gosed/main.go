package main

import "github.com/funtimecoding/soil/pkg/tool/gosed"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosed.Main(Version, GitHash, BuildDate)
}
