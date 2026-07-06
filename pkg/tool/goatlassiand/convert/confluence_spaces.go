package convert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func ConfluenceSpaces(v []*space.Space) []*server.ConfluenceSpace {
	result := make([]*server.ConfluenceSpace, 0, len(v))

	for _, s := range v {
		p := &server.ConfluenceSpace{
			Identifier: s.Identifier,
			Key:        s.Raw.Key,
			Name:       s.Name,
		}

		if s.Raw.Type != "" {
			p.Type = &s.Raw.Type
		}

		if s.Raw.Description != "" {
			p.Description = &s.Raw.Description
		}

		if s.Raw.HomepageIdentifier != "" {
			p.HomepageIdentifier = &s.Raw.HomepageIdentifier
		}

		result = append(result, p)
	}

	return result
}
