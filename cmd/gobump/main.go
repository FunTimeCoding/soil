package main

import "github.com/funtimecoding/soil/pkg/tool/gobump"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gobump.Main(Version, GitHash, BuildDate)
}
