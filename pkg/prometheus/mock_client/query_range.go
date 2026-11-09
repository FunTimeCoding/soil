package mock_client

import (
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

func (c *Client) QueryRange(
	_ string,
	r v1.Range,
) (*query_result.Result, error) {
	c.record(r.Start, r.End)

	if c.Failure != nil {
		return nil, c.Failure
	}

	return query_result.New(model.Matrix{}, nil, nil), nil
}
