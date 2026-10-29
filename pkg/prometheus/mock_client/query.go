package mock_client

import (
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"github.com/prometheus/common/model"
	"time"
)

func (c *Client) Query(
	q string,
	_ time.Time,
) (*query_result.Result, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e, okay := c.failures[q]; okay {
		return nil, e
	}

	if result, okay := c.results[q]; okay {
		return result, nil
	}

	return query_result.New(model.Vector{}, nil, nil), nil
}
