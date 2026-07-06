package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/brave"
	"github.com/funtimecoding/soil/pkg/brave/bookmark/node"
	"github.com/funtimecoding/soil/pkg/brave/constant"
)

func BookmarkNode() {
	a := argument.NewSimple("bookmark-node")
	a.Integer(argument.Depth, 0, "")
	a.ParseSimple()
	depth := a.GetInteger(argument.Depth)
	directory := a.RequiredPositional(0, "DIRECTORY")
	b := brave.Bookmark(constant.DefaultProfile)
	f := constant.Format
	d := node.MustDirectoryByName(node.New(b.Root.Bar), directory)
	fmt.Printf("Root: %s\n", d.Format(f))

	if depth > 0 {
		node.WalkLevels(
			d,
			depth,
			func(o *node.Node) {
				fmt.Println(o.Format(f))
			},
		)
	} else {
		for _, l := range node.Collect(d) {
			fmt.Println(l.Format(f))
		}
	}

	for _, g := range node.GroupByDirectory(d) {
		fmt.Printf(
			"Group %s (%d)\n",
			g.Directory.Name,
			len(g.Links),
		)
	}
}
