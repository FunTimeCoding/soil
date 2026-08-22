package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) Query(
	q string,
	response any,
) error {
	_, e := c.client.GraphQL.Do(
		gitlab.GraphQLQuery{Query: q},
		response,
		gitlab.WithContext(c.context),
	)

	return wrapError(e)
}
