package main

import "github.com/funtimecoding/soil/pkg/tool/goproxmox"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goproxmox.Main(Version, GitHash, BuildDate)
}
