package main

import "github.com/funtimecoding/soil/pkg/tool/gosourced"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosourced.Main(Version, GitHash, BuildDate)
}
