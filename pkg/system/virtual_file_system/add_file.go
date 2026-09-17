package virtual_file_system

import "time"

func (s *System) AddFile(
	path string,
	content []byte,
	modTime time.Time,
) {
	s.files[path] = &File{
		Content: content,
		Size:    int64(len(content)),
		ModTime: modTime,
		Loaded:  true,
	}
}
