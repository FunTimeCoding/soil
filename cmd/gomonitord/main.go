package main

import "github.com/funtimecoding/soil/pkg/tool/gomonitord"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gomonitord.Main(Version, GitHash, BuildDate)
}
