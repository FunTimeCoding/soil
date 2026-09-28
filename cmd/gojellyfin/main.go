package main

import "github.com/funtimecoding/soil/pkg/tool/gojellyfin"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gojellyfin.Main(Version, GitHash, BuildDate)
}
