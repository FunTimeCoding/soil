package gitlab

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) Features() ([]*gitlab.Feature, error) {
	result, _, e := c.client.Features.ListFeatures()

	return result, wrapError(e)
}
