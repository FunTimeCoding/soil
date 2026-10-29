package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (c *Client) GetList(
	_ context.Context,
	p *client.GetListParams,
) (*client.ListOutcome, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	var found []client.SearchResult

	for _, d := range c.inCollection(p.Collection) {
		if matches(d, p.Metadata) {
			found = append(found, d)
		}
	}

	if p.Offset != nil {
		found = found[min(*p.Offset, len(found)):]
	}

	if p.Limit != nil {
		found = found[:min(*p.Limit, len(found))]
	}

	return &client.ListOutcome{Results: found}, nil
}
