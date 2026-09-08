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
		detail, f := c.generated.GetSessionDetailWithResponse(
			context.Background(),
			s.Identifier,
		)

		if f != nil {
			return nil, f
		}

		if detail.JSON200 == nil {
			continue
		}

		name := ""

		if s.Name != nil {
			name = *s.Name
		}

		timestamp, g := time.Parse(time.RFC3339Nano, s.Timestamp)

		if g != nil {
			timestamp = time.Time{}
		}

		labels := map[string]string{}

		if detail.JSON200.Labels != nil {
			for _, l := range *detail.JSON200.Labels {
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
