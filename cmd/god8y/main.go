package main

import "github.com/funtimecoding/soil/pkg/tool/god8y"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	god8y.Main(Version, GitHash, BuildDate)
}
