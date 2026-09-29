package service

import "github.com/funtimecoding/soil/pkg/linkace/link"

func (s *Service) EditLink(
	identifier int,
	o EditLinkOptions,
) (*link.Link, error) {
	existing, e := s.client.LinkByIdentifier(identifier)

	if e != nil {
		return nil, e
	}

	patch := map[string]any{
		"url":   existing.Link,
		"lists": existing.ListIdentifiers,
		"tags":  existing.TagIdentifiers,
	}

	if o.Name != "" {
		patch["title"] = o.Name
	}

	if o.Link != "" {
		patch["url"] = o.Link
	}

	if o.Description != "" {
		patch["description"] = o.Description
	}

	if o.Tags != nil {
		patch["tags"] = o.Tags
	} else {
		tags := existing.TagIdentifiers

		if len(o.AddTags) > 0 {
			for _, name := range o.AddTags {
				identifier, f := s.resolveTag(name)

				if f != nil {
					return nil, f
				}

				if !containsInteger(tags, identifier) {
					tags = append(tags, identifier)
				}
			}
		}

		if len(o.RemoveTags) > 0 {
			for _, name := range o.RemoveTags {
				identifier, f := s.resolveTag(name)

				if f != nil {
					return nil, f
				}

				tags = removeInteger(tags, identifier)
			}
		}

		patch["tags"] = tags
	}

	if o.Lists != nil {
		var identifiers []int

		for _, name := range o.Lists {
			identifier, f := s.ResolveList(name)

			if f != nil {
				return nil, f
			}

			identifiers = append(identifiers, identifier)
		}

		patch["lists"] = identifiers
	} else {
		lists := existing.ListIdentifiers

		if len(o.AddLists) > 0 {
			for _, name := range o.AddLists {
				identifier, f := s.ResolveList(name)

				if f != nil {
					return nil, f
				}

				if !containsInteger(lists, identifier) {
					lists = append(lists, identifier)
				}
			}
		}

		if len(o.RemoveLists) > 0 {
			for _, name := range o.RemoveLists {
				identifier, f := s.ResolveList(name)

				if f != nil {
					return nil, f
				}

				lists = removeInteger(lists, identifier)
			}
		}

		patch["lists"] = lists
	}

	return s.client.UpdateLink(identifier, patch)
}
