package service

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"

func (s *Service) forget(k *watched_request.Key) error {
	remaining, e := s.store.ByRequest(k)

	if e != nil {
		return e
	}

	identifiers := make([]uint, 0, len(remaining))

	for _, v := range remaining {
		identifiers = append(identifiers, v.Identifier)
	}

	s.pending.Retain(k.String(), identifiers)

	return s.store.DropOrphan(k)
}
