package main

import "github.com/funtimecoding/soil/pkg/tool/goterraformd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goterraformd.Main(Version, GitHash, BuildDate)
}
