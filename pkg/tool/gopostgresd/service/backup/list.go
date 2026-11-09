package backup

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/types/backup"

func (s *Service) List(instance string) ([]*backup.Backup, error) {
	if _, e := s.find(instance); e != nil {
		return nil, e
	}

	return s.store.List(instance)
}
