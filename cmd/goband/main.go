package main

import "github.com/funtimecoding/soil/pkg/tool/goband"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goband.Main(Version, GitHash, BuildDate)
}
