package main

import "github.com/funtimecoding/soil/pkg/tool/goatlassian"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goatlassian.Main(Version, GitHash, BuildDate)
}
