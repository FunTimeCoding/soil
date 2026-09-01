package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/brave"
	"github.com/funtimecoding/soil/pkg/brave/bookmark/node"
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func BookmarkNode() {
	a := argument.NewSimple("bookmark-node")
	a.Integer(argumentConstant.Depth, 0, "")
	a.ParseSimple()
	depth := a.GetInteger(argumentConstant.Depth)
	directory := a.RequiredPositional(0, "DIRECTORY")
	b := brave.Bookmark(constant.DefaultProfile)
	f := constant.Format
	d := node.MustDirectoryByName(node.New(b.Root.Bar), directory)
	console.Format("Root: %s\n", d.Format(f))

	if depth > 0 {
		node.WalkLevels(
			d,
			depth,
			func(o *node.Node) {
				console.Line(o.Format(f))
			},
		)
	} else {
		for _, l := range node.Collect(d) {
			console.Line(l.Format(f))
		}
	}

	for _, g := range node.GroupByDirectory(d) {
		console.Format("Group %s (%d)\n", g.Directory.Name, len(g.Links))
	}
}
