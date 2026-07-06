package main

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gomemoryd.Main(Version, GitHash, BuildDate)
}
