package main

import "github.com/funtimecoding/soil/pkg/tool/gocertificated"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gocertificated.Main(Version, GitHash, BuildDate)
}
