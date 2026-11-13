package virtual_file_system

import "github.com/funtimecoding/soil/pkg/system/virtual_file_system/file"

func (s *System) put(
	path string,
	f *file.File,
) {
	if _, present := s.files[path]; !present {
		s.sorted = nil
	}

	s.files[path] = f
}
