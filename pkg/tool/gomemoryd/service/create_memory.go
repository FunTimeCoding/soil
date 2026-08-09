package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func (s *Service) CreateMemory(o *save_option.Option) (*store.Memory, error) {
	if o.Scope == constant.AllScope || o.Scope == constant.DefaultScope {
		return nil, fmt.Errorf("%w: %s", constant.ErrorReservedScope, o.Scope)
	}

	if o.Type == "" {
		o.Type = "feedback"
	}

	identifier, e := s.store.CreateMemory(o)

	if e != nil {
		return nil, e
	}

	m, e := s.store.GetMemory(identifier)

	if e != nil {
		return nil, e
	}

	if e = s.syncIndex(m); e != nil {
		return nil, e
	}

	return m, nil
}
