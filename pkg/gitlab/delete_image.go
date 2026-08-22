package gitlab

func (c *Client) DeleteImage(
	project int64,
	repository int64,
	tag string,
) error {
	_, e := c.client.ContainerRegistry.DeleteRegistryRepositoryTag(
		project,
		repository,
		tag,
	)

	return wrapError(e)
}
