package main

import "github.com/funtimecoding/soil/pkg/tool/gotechnitium"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gotechnitium.Main(Version, GitHash, BuildDate)
}
