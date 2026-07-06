package main

import "github.com/funtimecoding/soil/pkg/tool/gov11y"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gov11y.Main(Version, GitHash, BuildDate)
}
