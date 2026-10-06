package web

import (
	"context"
	"net/http"
)

func DeferredBearerEditor(
	read func(string) string,
	name string,
) func(
	context.Context,
	*http.Request,
) error {
	return func(
		_ context.Context,
		q *http.Request,
	) error {
		Bearer(q, read(name))

		return nil
	}
}
