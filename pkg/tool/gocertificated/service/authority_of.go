package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/authority"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
)

func (s *Service) authorityOf(name string) (*authority.Authority, error) {
	r, e := s.store.Authority(name)

	if e != nil {
		return nil, e
	}

	if r == nil {
		return nil, not_found.New(constant.AuthorityParameter, name)
	}

	return authority.New(r.Material()), nil
}
