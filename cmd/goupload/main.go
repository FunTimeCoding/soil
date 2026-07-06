package main

import "github.com/funtimecoding/soil/pkg/tool/goupload"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goupload.Main(Version, GitHash, BuildDate)
}
