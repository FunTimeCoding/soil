package brave

import (
	"github.com/funtimecoding/soil/pkg/brave/bookmark/file"
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func Bookmark(profile string) *file.Bookmark {
	return file.Parse(
		system.ReadFile(
			MustProfileByName(profile).Path,
			constant.BookmarksFile,
		),
	)
}
