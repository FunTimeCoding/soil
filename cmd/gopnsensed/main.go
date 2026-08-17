package main

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gopnsensed.Main(Version, GitHash, BuildDate)
}
