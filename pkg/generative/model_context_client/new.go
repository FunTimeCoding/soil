package model_context_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client/bearer"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"net/http"
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
			Name:    soil.TestClient,
			Version: soil.DefaultVersion,
		},
		nil,
	).Connect(
		x,
		&mcp.StreamableClientTransport{
			Endpoint: locator.New(webConstant.Localhost).Insecure().Port(port).Path(
				webConstant.ModelContextPath,
			).String(),
			HTTPClient: &http.Client{Transport: bearer.Transport{}},
		},
		nil,
	)
	assert.FatalOnError(t, e)
	result := &Client{t: t, context: x, session: s}
	t.Cleanup(result.Close)

	return result
}
