package main

import "github.com/funtimecoding/soil/pkg/tool/godownload"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	godownload.Main(Version, GitHash, BuildDate)
}
