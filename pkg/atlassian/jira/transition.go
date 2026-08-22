package jira

func (c *Client) Transition(
	key string,
	transitionIdentifier string,
) error {
	_, e := c.client.Issue.DoTransitionWithContext(
		c.context,
		key,
		transitionIdentifier,
	)

	return wrapError(e)
}
