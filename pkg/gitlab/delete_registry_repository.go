package gitlab

func (c *Client) DeleteRegistryRepository(
	project int64,
	repository int64,
) error {
	_, e := c.client.ContainerRegistry.DeleteRegistryRepository(
		project,
		repository,
	)

	return wrapError(e)
}
