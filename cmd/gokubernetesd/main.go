package main

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gokubernetesd.Main(Version, GitHash, BuildDate)
}
