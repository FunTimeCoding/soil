package mock_client

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"

func (c *Client) SilenceByRule(_ string) (*silence.Silence, error) {
	return nil, c.failure
}
