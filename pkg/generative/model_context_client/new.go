package model_context_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/location"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"testing"
)

func New(
	t *testing.T,
	port int,
) *Client {
	t.Helper()
	x := context.Background()
	s, e := mcp.NewClient(
		&mcp.Implementation{
			Name:    constant.TestClient,
			Version: constant.DefaultVersion,
		},
		nil,
	).Connect(
		x,
		&mcp.StreamableClientTransport{
			Endpoint: locator.New(
				web.Localhost,
			).Insecure().Port(port).Path(
				location.ModelContext,
			).String(),
		},
		nil,
	)
	assert.FatalOnError(t, e)

	return &Client{t: t, context: x, session: s}
}
