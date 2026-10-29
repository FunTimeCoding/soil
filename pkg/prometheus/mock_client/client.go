package mock_client

import "github.com/funtimecoding/soil/pkg/prometheus/query_result"

type Client struct {
	results  map[string]*query_result.Result
	floats   map[string]float64
	failures map[string]error
	Failure  error
	Closed   int
}
