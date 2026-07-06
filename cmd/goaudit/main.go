package main

import "github.com/funtimecoding/soil/pkg/tool/goaudit"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goaudit.Main(Version, GitHash, BuildDate)
}
