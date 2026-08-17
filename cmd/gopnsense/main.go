package main

import "github.com/funtimecoding/soil/pkg/tool/gopnsense"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gopnsense.Main(Version, GitHash, BuildDate)
}
