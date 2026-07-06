package main

import "github.com/funtimecoding/soil/pkg/tool/goversion"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goversion.Main(Version, GitHash, BuildDate)
}
