package main

import "github.com/funtimecoding/soil/pkg/tool/gosublimed"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosublimed.Main(Version, GitHash, BuildDate)
}
