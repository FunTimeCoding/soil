package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) UpdateComment(
	key string,
	comment *jira.Comment,
) error {
	_, _, e := c.client.Issue.UpdateCommentWithContext(c.context, key, comment)

	return wrapError(e)
}
