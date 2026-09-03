package web

import (
	"context"
	"net/http"
)

func BearerEditor(token string) func(context.Context, *http.Request) error {
	return func(_ context.Context, q *http.Request) error {
		Bearer(q, token)

		return nil
	}
}
