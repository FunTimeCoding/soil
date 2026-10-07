package virtual_file_system

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system/file"
	"time"
)

func (s *System) AddMetadata(
	path string,
	size int64,
	modTime time.Time,
) {
	s.files[path] = file.New(size, modTime)
}
