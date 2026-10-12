package downstream

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func (c *Client) Trigger(changes []string) error {
	x := context.Background()
	s, e := mcp.NewClient(
		&mcp.Implementation{
			Name:    constant.DownstreamClient,
			Version: soil.DefaultVersion,
		},
		nil,
	).Connect(
		x,
		&mcp.StreamableClientTransport{
			Endpoint: locator.New(c.host).Insecure().Port(c.port).Path(
				webConstant.ModelContextPath,
			).String(),
		},
		nil,
	)

	if e != nil {
		return e
	}

	defer errors.PanicClose(s)
	_, f := s.CallTool(
		x,
		&mcp.CallToolParams{
			Name: constant.DownstreamTool,
			Arguments: map[string]any{
				constant.DownstreamUpdate:  true,
				constant.DownstreamChanges: changes,
			},
		},
	)

	return f
}
