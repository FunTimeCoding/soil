package service

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/types/discovered_file"

func (s *Service) Sync(files []discovered_file.File) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var paths []string

	for _, f := range files {
		paths = append(paths, f.Path)
		s.store.UpsertSeed(
			f.Name,
			f.Path,
			f.ContentHash,
			f.Content,
			f.ModifiedAt,
		)
	}

	s.store.RemoveMissing(paths)
	s.store.Compact()
	s.notifier.Notify()
}
