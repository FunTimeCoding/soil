package goreach

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/system"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func resolveSession(query string) string {
	c := goclauded.NewEnvironment()
	defer c.Close()
	r, matches, e := c.GetResolve(
		context.Background(),
		&client.GetResolveParams{Query: query},
	)

	if not_found.Is(e) {
		system.Exitf(1, "session not found: %s\n", query)
	}

	if e != nil {
		system.Exitf(1, "%s\n", e)
	}

	if len(matches) > 0 {
		system.Exitf(
			1,
			"ambiguous: %q matches %d sessions - pass its identifier\n",
			query,
			len(matches),
		)
	}

	return r.Identifier
}
