package sentry

import "fmt"

func (c *Client) DeleteIssue(
	organization string,
	identifier string,
) error {
	return c.basic.Delete(
		fmt.Sprintf("organizations/%s/issues/%s", organization, identifier),
	)
}
