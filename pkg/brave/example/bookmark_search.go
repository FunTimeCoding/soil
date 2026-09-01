package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/brave"
	"github.com/funtimecoding/soil/pkg/brave/bookmark/node"
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func BookmarkSearch() {
	a := argument.NewSimple("bookmark-search")
	a.String(argumentConstant.Type, constant.DirectoryType, "bookmark type")
	a.ParseSimple()
	search := a.RequiredPositional(0, "NAME")
	bookmarkType := a.GetString(argumentConstant.Type)
	b := brave.Bookmark(constant.DefaultProfile)
	f := constant.Format
	root := node.New(b.Root.Bar)
	results := node.FindAllByNameAndType(root, search, bookmarkType)
	node.SetParents(root)

	if len(results) == 0 {
		console.Line("No results")

		return
	}

	for _, n := range results {
		console.Line(n.Format(f))
		console.Format("  %s\n", n.FormatPath())
	}
}
