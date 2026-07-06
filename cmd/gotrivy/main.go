package main

import "github.com/funtimecoding/soil/pkg/tool/gotrivy"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gotrivy.Main(Version, GitHash, BuildDate)
}
