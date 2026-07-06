package main

import "github.com/funtimecoding/soil/pkg/tool/gogenie"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gogenie.Main(Version, GitHash, BuildDate)
}
