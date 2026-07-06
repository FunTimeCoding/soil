package main

import "github.com/funtimecoding/soil/pkg/tool/gochromed"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gochromed.Main(Version, GitHash, BuildDate)
}
