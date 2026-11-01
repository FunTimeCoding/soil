package mock_client

import "bytes"

func (c *Client) ArtifactFile(_ int64, _ int64, _ string) (*bytes.Reader, error) {
	return nil, c.failure
}
