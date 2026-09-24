package file

import "github.com/funtimecoding/soil/pkg/measure/count"

func New(
	path string,
	languageName string,
	c *count.Count,
) *File {
	return &File{Path: path, Language: languageName, Count: c}
}
