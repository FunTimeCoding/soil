package virtual_file_system

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system/file"
	"time"
)

func (s *System) Write(
	path string,
	content []byte,
) {
	s.checkLimits(len(content))
	f := file.New(int64(len(content)), time.Now())
	f.Content = content
	f.Loaded = true
	s.files[path] = f
	s.written[path] = true
}
