package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func ReadMessages(
	c *client.ClientWithResponses,
	identifiers []int,
) (string, error) {
	response, e := c.GetMessagesWithResponse(
		context.Background(),
		&client.GetMessagesParams{Identifier: identifiers},
	)

	if e != nil {
		return "", e
	}

	if response.JSON200 == nil {
		return "", fmt.Errorf(
			"%s: %s",
			response.Status(),
			string(response.Body),
		)
	}

	var lines []string

	for _, m := range response.JSON200.Messages {
		lines = append(
			lines,
			fmt.Sprintf(
				"[Message %d from %s to %s at %s]",
				m.Identifier,
				m.From,
				m.To,
				m.Timestamp,
			),
			m.Body,
			"",
		)
	}

	for _, i := range response.JSON200.Missing {
		lines = append(lines, fmt.Sprintf("No message with identifier: %d", i))
	}

	return join.NewLine(lines), nil
}
