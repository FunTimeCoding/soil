package main

import "github.com/funtimecoding/soil/pkg/tool/goansibled"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goansibled.Main(Version, GitHash, BuildDate)
}
