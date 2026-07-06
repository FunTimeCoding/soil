package main

import "github.com/funtimecoding/soil/pkg/tool/goalertmanagerd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	goalertmanagerd.Main(Version, GitHash, BuildDate)
}
