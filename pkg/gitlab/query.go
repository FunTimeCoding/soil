package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Query(
	x context.Context,
	q string,
	response any,
) error {
	_, e := c.client.GraphQL.Do(
		gitlab.GraphQLQuery{Query: q},
		response,
		gitlab.WithContext(x),
	)

	return wrapError(e)
}
