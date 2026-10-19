package basic

func (c *Client) Rejecting() *Client {
	return &Client{requester: c.requester, refusal: rejected}
}
