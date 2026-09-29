package service

import "github.com/funtimecoding/soil/pkg/linkace/link"

func (s *Service) AppendTag(
	linkIdentifier int,
	tagName string,
) (*link.Link, error) {
	tagIdentifier, e := s.resolveTag(tagName)

	if e != nil {
		return nil, e
	}

	existing, f := s.client.LinkByIdentifier(linkIdentifier)

	if f != nil {
		return nil, f
	}

	for _, v := range existing.TagIdentifiers {
		if v == tagIdentifier {
			return existing, nil
		}
	}

	ids := append(existing.TagIdentifiers, tagIdentifier)

	return s.client.UpdateLink(
		linkIdentifier,
		map[string]any{
			"url":   existing.Link,
			"lists": existing.ListIdentifiers,
			"tags":  ids,
		},
	)
}
