package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/brave"
	"github.com/funtimecoding/soil/pkg/brave/bookmark/file"
	"github.com/funtimecoding/soil/pkg/brave/constant"
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
		fmt.Println(n.Format(f))
	}
}
