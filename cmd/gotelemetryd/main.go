package main

import "github.com/funtimecoding/soil/pkg/tool/gotelemetryd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gotelemetryd.Main(Version, GitHash, BuildDate)
}
