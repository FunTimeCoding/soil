package jira

func (c *Client) DeleteComment(
	key string,
	identifier string,
) error {
	return wrapError(
		c.client.Issue.DeleteCommentWithContext(c.context, key, identifier),
	)
}
