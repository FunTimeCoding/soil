package file

import "time"

func New(
	size int64,
	modTime time.Time,
) *File {
	return &File{Size: size, ModTime: modTime}
}
