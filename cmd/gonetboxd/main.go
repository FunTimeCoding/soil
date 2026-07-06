package main

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gonetboxd.Main(Version, GitHash, BuildDate)
}
