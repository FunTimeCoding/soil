package main

import "github.com/funtimecoding/soil/pkg/tool/gopackagedeb"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gopackagedeb.Main(Version, GitHash, BuildDate)
}
