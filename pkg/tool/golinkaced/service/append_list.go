package service

import "github.com/funtimecoding/soil/pkg/linkace/link"

func (s *Service) AppendList(
	linkIdentifier int,
	listName string,
) (*link.Link, error) {
	listIdentifier, e := s.ResolveList(listName)

	if e != nil {
		return nil, e
	}

	existing, f := s.client.LinkByIdentifier(linkIdentifier)

	if f != nil {
		return nil, f
	}

	for _, v := range existing.ListIdentifiers {
		if v == listIdentifier {
			return existing, nil
		}
	}

	ids := append(existing.ListIdentifiers, listIdentifier)

	return s.client.UpdateLink(
		linkIdentifier,
		map[string]any{
			"url":   existing.Link,
			"lists": ids,
			"tags":  existing.TagIdentifiers,
		},
	)
}
