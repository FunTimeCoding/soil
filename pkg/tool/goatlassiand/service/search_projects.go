package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/project_match"
	"strings"
)

func (s *Service) SearchProjects(
	query string,
	limit int,
	detail bool,
) ([]*project_match.Match, error) {
	return guarded(
		s.jiraGuard,
		func() ([]*project_match.Match, error) {
			projects, e := s.jira.Projects()

			if e != nil {
				return nil, e
			}

			var result []*project_match.Match

			for _, p := range *projects {
				if len(result) >= limit {
					break
				}

				if query != "" &&
					!strings.EqualFold(p.Key, query) &&
					!strings.Contains(
						strings.ToLower(p.Name),
						strings.ToLower(query),
					) {
					continue
				}

				m := project_match.New(p.Key, p.Name, p.ProjectTypeKey)

				if detail {
					d, f := s.jira.Project(p.Key)

					if f != nil {
						return nil, f
					}

					m.Detail = d
				}

				result = append(result, m)
			}

			return result, nil
		},
	)
}
