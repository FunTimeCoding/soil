package model_context_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"net/http"
	"testing"
)

func New(
	t *testing.T,
	port int,
	base http.RoundTripper,
) *Client {
	t.Helper()
	x := context.Background()
	s, e := Session(x, port, base)
	assert.FatalOnError(t, e)
	result := &Client{t: t, context: x, session: s}
	t.Cleanup(result.Close)

	return result
}
