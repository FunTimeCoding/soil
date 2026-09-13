package main

import "github.com/funtimecoding/soil/pkg/tool/godirectory"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	godirectory.Main(Version, GitHash, BuildDate)
}
