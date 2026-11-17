package mock_client

import (
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		results:  make(map[string]*query_result.Result),
		floats:   make(map[string]float64),
		failures: make(map[string]error),
	}
	t.Cleanup(result.Close)

	return result
}
