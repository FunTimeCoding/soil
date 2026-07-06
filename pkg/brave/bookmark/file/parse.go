package file

import "github.com/funtimecoding/soil/pkg/notation"

func Parse(s string) *Bookmark {
	result := &Bookmark{}
	notation.MustDecode(s, result, false)

	return result
}
