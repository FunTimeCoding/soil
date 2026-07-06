package main

import "github.com/funtimecoding/soil/pkg/tool/gopostgres"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gopostgres.Main(Version, GitHash, BuildDate)
}
