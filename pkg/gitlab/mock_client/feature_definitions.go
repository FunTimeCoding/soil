package mock_client

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) FeatureDefinitions() ([]*gitlab.FeatureDefinition, error) {
	return nil, c.failure
}
