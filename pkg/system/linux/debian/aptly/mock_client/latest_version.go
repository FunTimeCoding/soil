package mock_client

func (c *Client) LatestVersion(
	repository string,
	name string,
) (string, error) {
	versions, e := c.Versions(repository, name)

	if e != nil {
		return "", e
	}

	if len(versions) == 0 {
		return "", nil
	}

	return versions[0], nil
}
