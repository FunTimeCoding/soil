package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) Organization(slug string) (*response.Organization, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result := response.NewOrganization()
	result.Slug = slug

	return result, nil
}
