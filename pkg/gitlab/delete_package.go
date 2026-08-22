package gitlab

func (c *Client) DeletePackage(
	project int64,
	repository int64,
) error {
	_, e := c.client.Packages.DeleteProjectPackage(project, repository)

	return wrapError(e)
}
