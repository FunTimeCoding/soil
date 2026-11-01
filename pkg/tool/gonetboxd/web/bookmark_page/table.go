package bookmark_page

import (
	"github.com/funtimecoding/soil/pkg/netbox/bookmark"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
)

func Table(bookmarks []*bookmark.Bookmark) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(bookmarks))

	for _, b := range bookmarks {
		rows = append(rows, row(b))
	}

	return table.New("Name", "Type", "Bookmarked").
		Rows(rows).
		Empty(constant.BookmarkNone).
		Node()
}
