package page_file

import "github.com/funtimecoding/soil/pkg/notation"

func Decode(s string) *File {
	var result *File
	notation.MustDecode(s, &result, false)

	return result
}
