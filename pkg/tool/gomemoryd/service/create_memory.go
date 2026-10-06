package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func (s *Service) CreateMemory(o *save_option.Option) (*record.Memory, error) {
	if o.Scope == constant.AllScope || o.Scope == constant.DefaultScope {
		return nil, validation.New("scope name is reserved: %s", o.Scope)
	}

	if o.Type == "" {
		o.Type = "feedback"
	}

	o.Metadata = applyBase(o.Metadata, o.Base)
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
