package main

import "github.com/funtimecoding/soil/pkg/tool/golinkace"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	golinkace.Main(Version, GitHash, BuildDate)
}
