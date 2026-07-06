package main

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goatlassiand.Main(Version, GitHash, BuildDate)
}
