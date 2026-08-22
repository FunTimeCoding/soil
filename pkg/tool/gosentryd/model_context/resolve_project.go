package model_context

import "github.com/funtimecoding/soil/pkg/errors/not_found"

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

	return "", not_found.New("project", slug)
}
