package main

import "github.com/funtimecoding/soil/pkg/tool/goghpr"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goghpr.Main(Version, GitHash, BuildDate)
}
