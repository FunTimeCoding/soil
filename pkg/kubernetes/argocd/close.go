package argocd

func (c *Client) Close() {
	c.requester.Close()
}
