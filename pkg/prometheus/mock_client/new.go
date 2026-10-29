package mock_client

import "github.com/funtimecoding/soil/pkg/prometheus/query_result"

func New() *Client {
	return &Client{
		results:  make(map[string]*query_result.Result),
		floats:   make(map[string]float64),
		failures: make(map[string]error),
	}
}
