package mock_client

import (
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"time"
)

func (c *Client) QueryExemplar(
	_ string,
	_ time.Time,
	_ time.Time,
) ([]v1.ExemplarQueryResult, error) {
	return nil, c.Failure
}
