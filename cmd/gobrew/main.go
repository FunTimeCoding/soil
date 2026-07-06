package main

import "github.com/funtimecoding/soil/pkg/tool/gobrew"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gobrew.Main(Version, GitHash, BuildDate)
}
