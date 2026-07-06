package main

import "github.com/funtimecoding/soil/pkg/tool/goname"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goname.Main(Version, GitHash, BuildDate)
}
