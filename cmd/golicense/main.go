package main

import "github.com/funtimecoding/soil/pkg/tool/golicense"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	golicense.Main(Version, GitHash, BuildDate)
}
