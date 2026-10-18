package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (s *Service) ContextPanel(
	x context.Context,
	sessionIdentifier string,
) string {
	var statistics *client.Statistics
	var relations []client.Relation

	if e := s.memoryGuard.Do(
		x,
		func(y context.Context) error {
			var f error

			if statistics, f = s.memory.Statistics(y); f != nil {
				return f
			}

			relations, f = s.memory.Relations(y)

			return f
		},
	); e != nil {
		return e.Error()
	}

	if statistics == nil {
		return ""
	}

	loaded := s.loadedIdentifiers(sessionIdentifier)
	var lines []string
	lines = append(lines, scopeLine(statistics, len(loaded)))

	if tags := tagLine(statistics, loaded); tags != "" {
		lines = append(lines, tags)
	}

	for _, one := range frontier(relations, loaded) {
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
