package main

import "github.com/funtimecoding/soil/pkg/tool/gorenovate"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gorenovate.Main(Version, GitHash, BuildDate)
}
