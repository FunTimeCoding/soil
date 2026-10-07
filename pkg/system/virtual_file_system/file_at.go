package virtual_file_system

import "github.com/funtimecoding/soil/pkg/system/virtual_file_system/file"

func (s *System) FileAt(path string) *file.File {
	return s.files[path]
}
