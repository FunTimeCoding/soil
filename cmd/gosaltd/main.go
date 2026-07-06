package main

import "github.com/funtimecoding/soil/pkg/tool/gosaltd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosaltd.Main(Version, GitHash, BuildDate)
}
