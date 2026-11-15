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

	if name != "" {
		return name, nil
	}

	if clusters := s.service.Clusters(); len(clusters) == 1 {
		return clusters[0].Name(), nil
	}

	return "", not_selected.Format(
		"no cluster selected - use use_cluster first",
	)
}
