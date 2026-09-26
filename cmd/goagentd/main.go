package main

import "github.com/funtimecoding/soil/pkg/tool/goagentd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goagentd.Main(Version, GitHash, BuildDate)
}
