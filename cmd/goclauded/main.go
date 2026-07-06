package main

import "github.com/funtimecoding/soil/pkg/tool/goclauded"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goclauded.Main(Version, GitHash, BuildDate)
}
