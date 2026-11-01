package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"
	"time"
)

func (c *Client) RecentSessions(limit int) ([]*target.Target, error) {
	sessions, e := c.GetSessions(
		context.Background(),
		&client.GetSessionsParams{Limit: &limit},
	)

	if e != nil {
		return nil, e
	}

	result := []*target.Target{}

	for _, s := range sessions.Sessions {
		name := ""

		if s.Name != nil {
			name = *s.Name
		}

		lastSeen := time.Time{}

		if s.LastSeen != nil {
			if parsed, f := time.Parse(
				time.RFC3339Nano,
				*s.LastSeen,
			); f == nil {
				lastSeen = parsed
			}
		}

		labels := map[string]string{}

		if s.Labels != nil {
			for _, l := range *s.Labels {
				labels[l.Key] = l.Value
			}
		}

		t := target.New(s.Identifier, name, lastSeen, labels)
		t.LastPromptAt = parseMoment(s.LastPromptAt)
		t.LastTurnEndAt = parseMoment(s.LastTurnEndAt)
		t.ClosedAt = parseMoment(s.ClosedAt)
		t.TurnCount = s.TurnCount
		result = append(result, t)
	}

	return result, nil
}
