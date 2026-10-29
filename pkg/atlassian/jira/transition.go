package jira

func (c *Client) Transition(
	key string,
	transitionIdentifier string,
) error {
	r, e := c.client.Issue.DoTransitionWithContext(
		c.context,
		key,
		transitionIdentifier,
	)

	return WrapError(r, e)
}
