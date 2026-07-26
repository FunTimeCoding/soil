package main

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/keepass"
	"github.com/funtimecoding/soil/pkg/keepass/constant"
)

func main() {
	c := keepass.NewEnvironment()
	entry := c.Root().Groups[0].Groups[0].Entries[0]
	fmt.Println(entry.GetTitle())

	if false {
		fmt.Println(entry.GetPassword())
	}

	fmt.Printf("Entry: %+v\n", c.ByTitle(constant.DirectoryTitle))
}
