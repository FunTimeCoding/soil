package mock_client

import "bytes"

func (c *Client) Artifacts(_ int64, _ int64) (*bytes.Reader, error) {
	return nil, c.failure
}
