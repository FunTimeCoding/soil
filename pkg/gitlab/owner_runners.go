package gitlab

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) ownerRunners(o *gitlab.ListRunnersOptions) ([]*gitlab.Runner, *gitlab.Response, error) {
	return c.client.Runners.ListRunners(o)
}
