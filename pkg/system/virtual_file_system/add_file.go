package virtual_file_system

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system/file"
	"time"
)

func (s *System) AddFile(
	path string,
	content []byte,
	modTime time.Time,
) {
	f := file.New(int64(len(content)), modTime)
	f.Content = content
	f.Loaded = true
	s.files[path] = f
}
