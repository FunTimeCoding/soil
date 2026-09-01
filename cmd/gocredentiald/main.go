package main

import "github.com/funtimecoding/soil/pkg/tool/gocredentiald"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gocredentiald.Main(Version, GitHash, BuildDate)
}
