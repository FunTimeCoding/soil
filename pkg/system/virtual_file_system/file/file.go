package file

import "time"

type File struct {
	Content []byte
	Size    int64
	ModTime time.Time
	Loaded  bool
}
