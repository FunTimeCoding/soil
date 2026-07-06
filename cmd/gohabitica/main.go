package main

import "github.com/funtimecoding/soil/pkg/tool/gohabitica"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gohabitica.Main(Version, GitHash, BuildDate)
}
