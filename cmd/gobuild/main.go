package main

import "github.com/funtimecoding/soil/pkg/tool/gobuild"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gobuild.Main(Version, GitHash, BuildDate)
}
