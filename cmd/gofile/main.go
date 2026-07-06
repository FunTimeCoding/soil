package main

import "github.com/funtimecoding/soil/pkg/tool/gofile"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gofile.Main(Version, GitHash, BuildDate)
}
