package example

import (
	"github.com/funtimecoding/soil/pkg/brave"
	"github.com/funtimecoding/soil/pkg/brave/bookmark/file"
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func BookmarkFile() {
	b := brave.Bookmark(constant.DefaultProfile)
	f := constant.Format
	var all []*file.Node
	file.Walk(
		b.Root.Bar,
		func(n *file.Node) {
			all = append(all, n)
		},
	)

	for _, n := range all {
		console.Line(n.Format(f))
	}
}
