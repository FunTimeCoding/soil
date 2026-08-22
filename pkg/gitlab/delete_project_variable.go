package gitlab

func (c *Client) DeleteProjectVariable(
	project int64,
	key string,
) error {
	_, e := c.client.ProjectVariables.RemoveVariable(project, key, nil)

	return wrapError(e)
}
