package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) CreateTag(
	project int64,
	name string,
	reference string,
	message string,
) (*gitlab.Tag, error) {
	result, _, e := c.client.Tags.CreateTag(
		project,
		&gitlab.CreateTagOptions{
			TagName: new(name),
			Ref:     new(reference),
			Message: new(message),
		},
	)

	return result, wrapError(e)
}
