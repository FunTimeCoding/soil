package main

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goproxmoxd.Main(Version, GitHash, BuildDate)
}
