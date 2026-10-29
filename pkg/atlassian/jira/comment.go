package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Comment(
	key string,
	body string,
) error {
	_, r, e := c.client.Issue.AddCommentWithContext(
		c.context,
		key,
		&jira.Comment{Body: body},
	)

	return WrapError(r, e)
}
