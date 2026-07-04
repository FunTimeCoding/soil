package main

import "github.com/funtimecoding/go-library/pkg/tool/godashboardd"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	godashboardd.Main(Version, GitHash, BuildDate)
}
