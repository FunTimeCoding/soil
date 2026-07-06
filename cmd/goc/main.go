package main

import "github.com/funtimecoding/soil/pkg/tool/goc"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goc.Main(Version, GitHash, BuildDate)
}
