package main

import "github.com/funtimecoding/soil/pkg/tool/goraidparsed"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goraidparsed.Main(Version, GitHash, BuildDate)
}
