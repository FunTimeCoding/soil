package main

import "github.com/funtimecoding/soil/pkg/tool/gosentry"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosentry.Main(Version, GitHash, BuildDate)
}
