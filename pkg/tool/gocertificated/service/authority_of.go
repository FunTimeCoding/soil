package service

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"

func (s *Service) authorityOf(name string) (*authority.Authority, error) {
	r, e := s.store.Authority(name)

	if e != nil {
		return nil, e
	}

	return authority.New(r.Material()), nil
}
