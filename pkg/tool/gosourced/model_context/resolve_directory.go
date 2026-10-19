package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
)

func (s *Server) resolveDirectory(x context.Context) (string, error) {
	name, e := s.selection.Active(x)

	if e != nil {
		return "", e
	}

	if name == "" {
		return "", not_selected.Format(
			"no active module - call use_module first",
		)
	}

	m, okay := s.service.Module(name)

	if !okay {
		return "", not_found.New(constant.ModuleKind, name)
	}

	return m.Directory, nil
}
