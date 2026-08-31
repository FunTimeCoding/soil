package downstream

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	provision "github.com/funtimecoding/soil/pkg/provision/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func (c *Client) Trigger(changes []string) error {
	x := context.Background()
	s, e := mcp.NewClient(
		&mcp.Implementation{
			Name:    provision.DownstreamClient,
			Version: constant.DefaultVersion,
		},
		nil,
	).Connect(
		x,
		&mcp.StreamableClientTransport{
			Endpoint: locator.New(c.host).Insecure().Port(c.port).Path(
				web.ModelContextPath,
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
			Name: provision.DownstreamTool,
			Arguments: map[string]any{
				provision.DownstreamUpdate:  true,
				provision.DownstreamChanges: changes,
			},
		},
	)

	return f
}
