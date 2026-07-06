package main

import "github.com/funtimecoding/soil/pkg/tool/godockerhub"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	godockerhub.Main(Version, GitHash, BuildDate)
}
