package main

import "github.com/funtimecoding/soil/pkg/tool/gofix"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gofix.Main(Version, GitHash, BuildDate)
}
