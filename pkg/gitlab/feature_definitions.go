package gitlab

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) FeatureDefinitions() ([]*gitlab.FeatureDefinition, error) {
	result, _, e := c.client.Features.ListFeatureDefinitions()

	return result, wrapError(e)
}
