package main

import "github.com/funtimecoding/soil/pkg/tool/golinkaced"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	golinkaced.Main(Version, GitHash, BuildDate)
}
