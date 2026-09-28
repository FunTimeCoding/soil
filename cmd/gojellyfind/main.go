package main

import "github.com/funtimecoding/soil/pkg/tool/gojellyfind"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gojellyfind.Main(Version, GitHash, BuildDate)
}
