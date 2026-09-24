package main

import "github.com/funtimecoding/soil/pkg/tool/gocrap"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gocrap.Main(Version, GitHash, BuildDate)
}
