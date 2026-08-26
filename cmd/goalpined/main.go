package main

import "github.com/funtimecoding/soil/pkg/tool/goalpined"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goalpined.Main(Version, GitHash, BuildDate)
}
