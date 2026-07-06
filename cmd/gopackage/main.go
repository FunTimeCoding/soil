package main

import "github.com/funtimecoding/soil/pkg/tool/gopackage"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gopackage.Main(Version, GitHash, BuildDate)
}
