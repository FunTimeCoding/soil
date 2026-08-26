package main

import "github.com/funtimecoding/soil/pkg/tool/gopackageapk"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gopackageapk.Main(Version, GitHash, BuildDate)
}
