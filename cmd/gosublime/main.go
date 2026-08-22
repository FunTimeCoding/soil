package main

import "github.com/funtimecoding/soil/pkg/tool/gosublime"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosublime.Main(Version, GitHash, BuildDate)
}
