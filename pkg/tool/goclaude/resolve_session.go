package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
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
		c.Terminal().Exitln(not_found.New(goclaudedConstant.SessionKind, query))

		return ""
	}

	if e != nil {
		c.Terminal().Exitln(e)

		return ""
	}

	if len(matches) > 0 {
		lines := []string{
			fmt.Sprintf(
				"ambiguous: %q matches %d sessions",
				query,
				len(matches),
			),
		}

		for _, m := range matches {
			name := ""

			if m.Name != nil {
				name = *m.Name
			}

			display := "(discovered)"

			if m.Alias != nil && *m.Alias != "" {
				display = *m.Alias
			}

			lines = append(
				lines,
				fmt.Sprintf("  %s  %-7s  %s", m.Identifier[:8], name, display),
			)
		}

		c.Terminal().Exitln(validation.New("%s", join.NewLine(lines)))

		return ""
	}

	return r.Identifier
}
