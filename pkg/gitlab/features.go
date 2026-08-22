package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) Features() ([]*gitlab.Feature, error) {
	result, _, e := c.client.Features.ListFeatures()

	return result, wrapError(e)
}
