package main

import "github.com/funtimecoding/soil/pkg/tool/gomemory"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gomemory.Main(Version, GitHash, BuildDate)
}
