package model_context_client

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client/bearer"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"net/http"
)

func Session(
	x context.Context,
	port int,
	base http.RoundTripper,
) (*mcp.ClientSession, error) {
	return mcp.NewClient(
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
			HTTPClient: &http.Client{Transport: bearer.Transport{Base: base}},
		},
		nil,
	)
}
