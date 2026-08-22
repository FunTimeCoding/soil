package gitlab

func (c *Client) DeleteTag(
	project int64,
	name string,
) error {
	_, e := c.client.Tags.DeleteTag(project, name)

	return wrapError(e)
}
