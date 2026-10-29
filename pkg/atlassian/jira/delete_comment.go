package jira

func (c *Client) DeleteComment(
	key string,
	identifier string,
) error {
	return WrapError(
		nil,
		c.client.Issue.DeleteCommentWithContext(c.context, key, identifier),
	)
}
