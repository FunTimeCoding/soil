package main

import "github.com/funtimecoding/soil/pkg/tool/gomaintlog"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gomaintlog.Main(Version, GitHash, BuildDate)
}
