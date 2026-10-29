package mock_client

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"

func (c *Client) Silences(_ bool) ([]*silence.Silence, error) {
	return nil, c.failure
}
