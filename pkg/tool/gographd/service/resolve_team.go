package service

import (
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"strings"
)

func (s *Service) resolveTeam(name string) (string, error) {
	if name == "" {
		return "", validation.New("team is required")
	}

	teams, e := s.Teams()

	if e != nil {
		return "", e
	}

	var matches []string

	for _, t := range teams {
		if t.Identifier == name {
			return t.Identifier, nil
		}

		if strings.EqualFold(t.Name, name) {
			matches = append(matches, t.Identifier)
		}
	}

	switch len(matches) {
	case 0:
		return "", not_found.New("team", name)
	case 1:
		return matches[0], nil
	default:
		return "", ambiguous.Format(
			"team name matches %d teams, use the identifier: %s",
			len(matches),
			name,
		)
	}
}
