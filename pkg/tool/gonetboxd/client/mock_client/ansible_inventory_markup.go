package mock_client

import "context"

func (c *Client) AnsibleInventoryMarkup(_ context.Context, _ string) (string, error) {
	return "", c.Failure
}
