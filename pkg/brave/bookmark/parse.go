package bookmark

import "github.com/funtimecoding/soil/pkg/notation"

func Parse(s string) *File {
	result := &File{}
	notation.MustDecode(s, result, false)

	return result
}
