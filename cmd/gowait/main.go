package main

import "github.com/funtimecoding/soil/pkg/tool/gowait"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gowait.Main(Version, GitHash, BuildDate)
}
