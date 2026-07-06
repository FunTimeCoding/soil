package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/brave"
	"github.com/funtimecoding/soil/pkg/brave/bookmark"
	"github.com/funtimecoding/soil/pkg/brave/bookmark/node"
	"github.com/funtimecoding/soil/pkg/brave/constant"
)

func BookmarkSearch() {
	a := argument.NewSimple("bookmark-search")
	a.String(argument.Type, bookmark.DirectoryType, "bookmark type")
	a.ParseSimple()
	search := a.RequiredPositional(0, "NAME")
	bookmarkType := a.GetString(argument.Type)
	b := brave.Bookmark(constant.DefaultProfile)
	f := constant.Format
	root := node.New(b.Root.Bar)
	results := node.FindAllByNameAndType(root, search, bookmarkType)
	node.SetParents(root)

	if len(results) == 0 {
		fmt.Println("No results")

		return
	}

	for _, n := range results {
		fmt.Println(n.Format(f))
		fmt.Printf("  %s\n", n.FormatPath())
	}
}
