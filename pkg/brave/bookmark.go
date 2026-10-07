package brave

import (
	"github.com/funtimecoding/soil/pkg/brave/bookmark"
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func Bookmark(profile string) *bookmark.File {
	return bookmark.Parse(
		system.ReadFile(
			MustProfileByName(profile).Path,
			constant.BookmarksFile,
		),
	)
}
