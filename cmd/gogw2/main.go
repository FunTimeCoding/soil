package main

import "github.com/funtimecoding/soil/pkg/tool/gogw2"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gogw2.Main(Version, GitHash, BuildDate)
}
