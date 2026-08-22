package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) Compare(
	project int64,
	from, to string,
) (*gitlab.Compare, error) {
	result, _, e := c.client.Repositories.Compare(
		project,
		&gitlab.CompareOptions{From: new(from), To: new(to)},
	)

	return result, wrapError(e)
}
