package main

import "github.com/funtimecoding/soil/pkg/tool/gosecret"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosecret.Main(Version, GitHash, BuildDate)
}
