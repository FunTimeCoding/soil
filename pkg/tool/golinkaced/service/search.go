package service

import (
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/types/search_result"
	"strings"
)

func (s *Service) Search(
	query string,
	entityType string,
) (*search_result.Result, error) {
	result := search_result.New()
	lower := strings.ToLower(query)

	if entityType == "" || entityType == "link" {
		links, e := s.client.Search(query)

		if e != nil {
			return nil, e
		}

		result.Links = links
	}

	if entityType == "" || entityType == "list" {
		all, e := s.client.Lists()

		if e != nil {
			return nil, e
		}

		for _, l := range all {
			if strings.Contains(strings.ToLower(l.Name), lower) {
				result.Lists = append(result.Lists, l)
			}
		}
	}

	if entityType == "" || entityType == "tag" {
		all, e := s.client.Tags()

		if e != nil {
			return nil, e
		}

		for _, t := range all {
			if strings.Contains(strings.ToLower(t.Name), lower) {
				result.Tags = append(result.Tags, t)
			}
		}
	}

	return result, nil
}
