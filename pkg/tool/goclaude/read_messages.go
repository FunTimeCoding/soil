package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func ReadMessages(
	c face.Client,
	identifiers []int,
) (string, error) {
	r, e := c.GetMessages(
		context.Background(),
		&client.GetMessagesParams{Identifier: identifiers},
	)

	if e != nil {
		return "", e
	}

	var lines []string

	for _, m := range r.Messages {
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

	for _, i := range r.Missing {
		lines = append(lines, fmt.Sprintf("No message with identifier: %d", i))
	}

	return join.NewLine(lines), nil
}
