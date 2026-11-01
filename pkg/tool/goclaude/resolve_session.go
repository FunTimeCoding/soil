package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func resolveSession(
	c *command_context.Context,
	query string,
) string {
	r, matches, e := c.Client().GetResolve(
		context.Background(),
		&client.GetResolveParams{Query: query},
	)

	if not_found.Is(e) {
		return ""
	}

	if e != nil {
		c.Terminal().Exitln(e)

		return ""
	}

	if len(matches) > 0 {
		console.Format(
			"ambiguous: %q matches %d sessions\n",
			query,
			len(matches),
		)

		for _, m := range matches {
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

			console.Format("  %s  %-7s  %s\n", m.Identifier[:8], name, display)
		}

		return ""
	}

	return r.Identifier
}
