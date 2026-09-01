package main

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/keepass"
	"github.com/funtimecoding/soil/pkg/keepass/constant"
)

func main() {
	c := keepass.NewEnvironment()
	entry := c.Root().Groups[0].Groups[0].Entries[0]
	console.Line(entry.GetTitle())

	if false {
		console.Line(entry.GetPassword())
	}

	console.Format("Entry: %+v\n", c.ByTitle(constant.DirectoryTitle))
}
