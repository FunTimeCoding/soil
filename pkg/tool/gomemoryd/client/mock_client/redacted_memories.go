package mock_client

import "context"

func (c *Client) RedactedMemories(context.Context) (map[int64]bool, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Redacted, nil
}
