package service

import "github.com/funtimecoding/soil/pkg/linkace/link"

func (s *Service) RemoveList(
	linkIdentifier int,
	listName string,
) (*link.Link, error) {
	listIdentifier, e := s.ResolveList(listName)

	if e != nil {
		return nil, e
	}

	existing, f := s.linkByIdentifier(linkIdentifier)

	if f != nil {
		return nil, f
	}

	var ids []int

	for _, v := range existing.ListIdentifiers {
		if v != listIdentifier {
			ids = append(ids, v)
		}
	}

	return s.updateLink(
		linkIdentifier,
		map[string]any{
			"url":   existing.Link,
			"lists": ids,
			"tags":  existing.TagIdentifiers,
		},
	)
}
