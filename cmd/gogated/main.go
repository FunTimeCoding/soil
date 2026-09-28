package main

import "github.com/funtimecoding/soil/pkg/tool/gogated"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gogated.Main(Version, GitHash, BuildDate)
}
