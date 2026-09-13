package main

import "github.com/funtimecoding/soil/pkg/tool/godirectoryd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	godirectoryd.Main(Version, GitHash, BuildDate)
}
