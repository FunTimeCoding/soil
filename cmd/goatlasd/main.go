package main

import "github.com/funtimecoding/soil/pkg/tool/goatlasd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goatlasd.Main(Version, GitHash, BuildDate)
}
