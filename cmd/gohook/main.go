package main

import "github.com/funtimecoding/soil/pkg/tool/gohook"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gohook.Main(Version, GitHash, BuildDate)
}
