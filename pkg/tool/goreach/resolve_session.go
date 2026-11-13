package goreach

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func resolveSession(query string) (string, error) {
	c := goclauded.NewEnvironment()
	defer c.Close()
	r, matches, e := c.GetResolve(
		context.Background(),
		&client.GetResolveParams{Query: query},
	)

	if not_found.Is(e) {
		return "", not_found.New("session", query)
	}

	if e != nil {
		return "", e
	}

	if len(matches) > 0 {
		return "", ambiguous.Format(
			"ambiguous: %q matches %d sessions - pass its identifier",
			query,
			len(matches),
		)
	}

	return r.Identifier, nil
}
