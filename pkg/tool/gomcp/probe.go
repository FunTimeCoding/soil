package gomcp

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
)

func probe(
	locator string,
	token string,
) {
	var opts []transport.StreamableHTTPCOption

	if token != "" {
		opts = append(
			opts,
			transport.WithHTTPHeaders(
				map[string]string{
					"Authorization": fmt.Sprintf("Bearer %s", token),
				},
			),
		)
	}

	c, e := client.NewStreamableHttpClient(locator, opts...)
	errors.PanicOnError(e)
	x := context.Background()
	errors.PanicOnError(c.Start(x))
	_, e = c.Initialize(x, mcp.InitializeRequest{})
	errors.PanicOnError(e)
	fmt.Printf("Session ID: %s\n", c.GetTransport().GetSessionId())
	result, f := c.ListTools(x, mcp.ListToolsRequest{})
	errors.PanicOnError(f)

	for _, t := range result.Tools {
		fmt.Printf("%s: %s\n", t.Name, t.Description)
	}
}
