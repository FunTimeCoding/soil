package main

import "github.com/funtimecoding/soil/pkg/tool/gomcp"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gomcp.Main(Version, GitHash, BuildDate)
}
