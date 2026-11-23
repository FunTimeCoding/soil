package service

import (
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"strings"
)

func (s *Service) resolveChannel(
	team string,
	name string,
) (string, string, error) {
	if name == "" {
		return "", "", validation.New("channel is required")
	}

	t, e := s.resolveTeam(team)

	if e != nil {
		return "", "", e
	}

	channels, e := s.teamChannels(t)

	if e != nil {
		return "", "", e
	}

	var matches []string

	for _, h := range channels {
		if h.Identifier == name {
			return t, h.Identifier, nil
		}

		if strings.EqualFold(h.Name, name) {
			matches = append(matches, h.Identifier)
		}
	}

	switch len(matches) {
	case 0:
		return "", "", not_found.New("channel", name)
	case 1:
		return t, matches[0], nil
	default:
		return "", "", ambiguous.Format(
			"channel name matches %d channels, use the identifier: %s",
			len(matches),
			name,
		)
	}
}
