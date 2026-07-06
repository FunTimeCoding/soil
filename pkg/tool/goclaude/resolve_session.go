package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func resolveSession(
	c *client.ClientWithResponses,
	query string,
) string {
	response, e := c.GetResolveWithResponse(
		context.Background(),
		&client.GetResolveParams{Query: query},
	)

	if e != nil {
		return ""
	}

	if response.StatusCode() == 409 && response.JSON409 != nil {
		fmt.Printf(
			"ambiguous: %q matches %d sessions\n",
			query,
			len(response.JSON409.Matches),
		)

		for _, m := range response.JSON409.Matches {
			name := ""

			if m.Name != nil {
				name = *m.Name
			}

			alias := ""

			if m.Alias != nil {
				alias = *m.Alias
			}

			display := alias

			if display == "" {
				display = "(discovered)"
			}

			fmt.Printf(
				"  %s  %-7s  %s\n",
				m.Identifier[:8],
				name,
				display,
			)
		}

		return ""
	}

	if response.JSON200 == nil {
		return ""
	}

	return response.JSON200.Identifier
}
