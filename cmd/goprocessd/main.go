package main

import "github.com/funtimecoding/soil/pkg/tool/goprocessd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goprocessd.Main(Version, GitHash, BuildDate)
}
