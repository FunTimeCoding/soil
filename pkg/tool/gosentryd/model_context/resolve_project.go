package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
)

func (s *Server) resolveProject(slug string) (string, error) {
	if slug == "" {
		return "", nil
	}

	projects, e := s.client.OrganizationProjects(s.organization)

	if e != nil {
		return "", e
	}

	for _, p := range projects {
		if p.Slug == slug {
			return p.ID, nil
		}
	}

	return "", fmt.Errorf("%w: %s", constant.ErrorProjectNotFound, slug)
}
