package main

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gogitlabd.Main(Version, GitHash, BuildDate)
}
