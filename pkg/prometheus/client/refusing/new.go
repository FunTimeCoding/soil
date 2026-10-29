package refusing

import "github.com/prometheus/client_golang/api"

func New(c api.Client) *Client {
	return &Client{client: c}
}
