package main

import "github.com/funtimecoding/soil/pkg/tool/goloc"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goloc.Main(Version, GitHash, BuildDate)
}
