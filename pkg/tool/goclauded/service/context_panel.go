package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Service) ContextPanel(sessionIdentifier string) string {
	statistics := s.memory.Statistics()

	if statistics == nil {
		return ""
	}

	loaded := s.loadedIdentifiers(sessionIdentifier)
	var lines []string
	lines = append(lines, scopeLine(statistics, len(loaded)))

	if tags := tagLine(statistics, loaded); tags != "" {
		lines = append(lines, tags)
	}

	for _, one := range frontier(s.memory.Relations(), loaded) {
		lines = append(
			lines,
			fmt.Sprintf(
				"door     %d  %s   %s <- %s",
				one.Identifier,
				one.Name,
				one.Relation,
				one.Source,
			),
		)
	}

	return join.NewLine(lines)
}
