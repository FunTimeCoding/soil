package main

import "github.com/funtimecoding/soil/pkg/tool/gogithubd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gogithubd.Main(Version, GitHash, BuildDate)
}
