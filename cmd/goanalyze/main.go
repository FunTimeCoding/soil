package main

import "github.com/funtimecoding/soil/pkg/tool/goanalyze"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goanalyze.Main(Version, GitHash, BuildDate)
}
