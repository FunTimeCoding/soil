package main

import "github.com/funtimecoding/soil/pkg/tool/gosilence"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosilence.Main(Version, GitHash, BuildDate)
}
