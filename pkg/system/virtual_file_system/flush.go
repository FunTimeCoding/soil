package virtual_file_system

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func (s *System) Flush(directory string) {
	for _, m := range s.moves {
		from := filepath.Join(directory, m.From)
		to := filepath.Join(directory, m.To)
		system.EnsurePathExists(filepath.Dir(to))
		system.Move(from, to)
	}

	for path := range s.written {
		f := s.files[path]

		if f == nil || !f.Loaded {
			continue
		}

		full := filepath.Join(directory, path)
		system.EnsurePathExists(filepath.Dir(full))
		system.SaveBytes(full, f.Content)
	}

	for path := range s.deleted {
		system.RemoveFile(filepath.Join(directory, path))
	}

	s.moves = nil
	s.written = make(map[string]bool)
	s.deleted = make(map[string]bool)
}
