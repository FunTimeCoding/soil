package page_file

import "github.com/funtimecoding/soil/pkg/notation"

func (f *File) Encode() string {
	return notation.Encode(f, true)
}
