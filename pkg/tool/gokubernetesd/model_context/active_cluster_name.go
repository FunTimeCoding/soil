package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
)

func (s *Server) activeClusterName(x context.Context) (string, error) {
	name, e := s.selection.Active(x)

	if e != nil {
		return "", e
	}

	if name == "" {
		return "", not_selected.Format(
			"no cluster selected - use use_cluster first",
		)
	}

	return name, nil
}
