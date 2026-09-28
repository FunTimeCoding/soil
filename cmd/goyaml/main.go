package main

import "github.com/funtimecoding/soil/pkg/tool/goyaml"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goyaml.Main(Version, GitHash, BuildDate)
}
