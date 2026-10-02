package main

import "github.com/funtimecoding/soil/pkg/tool/goflow"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goflow.Main(Version, GitHash, BuildDate)
}
