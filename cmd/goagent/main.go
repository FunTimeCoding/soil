package main

import "github.com/funtimecoding/soil/pkg/tool/goagent"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goagent.Main(Version, GitHash, BuildDate)
}
