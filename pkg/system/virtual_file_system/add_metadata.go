package virtual_file_system

import "time"

func (s *System) AddMetadata(
	path string,
	size int64,
	modTime time.Time,
) {
	s.files[path] = &File{Size: size, ModTime: modTime}
}
