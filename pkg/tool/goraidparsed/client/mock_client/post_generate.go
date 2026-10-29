package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/client"
	"time"
)

func (c *Client) PostGenerate(
	_ context.Context,
	files []string,
	_ *time.Time,
) (client.GenerateResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Calls++

	if c.Failure != nil {
		return client.GenerateResponse{}, c.Failure
	}

	c.Generated = append(c.Generated, files)

	return client.GenerateResponse{Path: "report.html", TidCount: len(files)}, nil
}
