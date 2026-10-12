package openbao

func (c *Client) WithKubernetes(
	role string,
	tokenPath string,
) *Client {
	c.role = role
	c.tokenPath = tokenPath

	return c
}
