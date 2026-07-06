package main

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gopostgresd.Main(Version, GitHash, BuildDate)
}
