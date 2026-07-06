package main

import "github.com/funtimecoding/soil/pkg/tool/goqueryd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goqueryd.Main(Version, GitHash, BuildDate)
}
