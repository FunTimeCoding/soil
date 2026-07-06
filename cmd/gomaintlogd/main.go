package main

import "github.com/funtimecoding/soil/pkg/tool/gomaintlogd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gomaintlogd.Main(Version, GitHash, BuildDate)
}
