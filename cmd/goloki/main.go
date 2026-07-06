package main

import "github.com/funtimecoding/soil/pkg/tool/goloki"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goloki.Main(Version, GitHash, BuildDate)
}
