package connector

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"time"
)

func (c *Client) RecentSessions(limit int) ([]*Target, error) {
	sessions, e := c.generated.GetSessionsWithResponse(
		context.Background(),
		&client.GetSessionsParams{Limit: &limit},
	)

	if e != nil {
		return nil, e
	}

	if sessions.JSON200 == nil {
		return nil, unexpected.Format(
			"session list status %d",
			sessions.StatusCode(),
		)
	}

	result := []*Target{}

	for _, s := range sessions.JSON200.Sessions {
		name := ""

		if s.Name != nil {
			name = *s.Name
		}

		timestamp := time.Time{}

		if s.LastSeen != nil {
			if parsed, f := time.Parse(
				time.RFC3339Nano,
				*s.LastSeen,
			); f == nil {
				timestamp = parsed
			}
		}

		labels := map[string]string{}

		if s.Labels != nil {
			for _, l := range *s.Labels {
				labels[l.Key] = l.Value
			}
		}

		result = append(
			result,
			NewTarget(s.Identifier, name, timestamp, labels),
		)
	}

	return result, nil
}
