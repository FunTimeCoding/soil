package main

import "github.com/funtimecoding/soil/pkg/tool/gomonitor"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gomonitor.Main(Version, GitHash, BuildDate)
}
