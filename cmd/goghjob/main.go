package main

import "github.com/funtimecoding/soil/pkg/tool/goghjob"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goghjob.Main(Version, GitHash, BuildDate)
}
