package main

import "github.com/funtimecoding/soil/pkg/tool/godebian"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	godebian.Main(Version, GitHash, BuildDate)
}
