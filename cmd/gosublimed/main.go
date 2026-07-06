// Setup: ln -s $HOME/src/go-library/py/sublime_bridge "$HOME/Library/Application Support/Sublime Text/Packages/SublimeBridge"
package main

import "github.com/funtimecoding/soil/pkg/tool/gosublimed"

var (
	Version   string
	GitHash   string
	BuildDate string
)

func main() {
	gosublimed.Main(Version, GitHash, BuildDate)
}
