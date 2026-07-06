package main

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gomattermostd.Main(Version, GitHash, BuildDate)
}
