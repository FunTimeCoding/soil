package main

import "github.com/funtimecoding/soil/pkg/tool/goatlas"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goatlas.Main(Version, GitHash, BuildDate)
}
