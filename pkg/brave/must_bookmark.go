package brave

import (
	"github.com/funtimecoding/soil/pkg/brave/bookmark"
	"github.com/funtimecoding/soil/pkg/errors"
)

func MustBookmark(profile string) *bookmark.File {
	result, e := Bookmark(profile)
	errors.PanicOnError(e)

	return result
}
