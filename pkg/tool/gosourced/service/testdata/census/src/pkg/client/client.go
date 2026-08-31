package client

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c *Client) Ready() bool {
	return true
}

func (c *Client) Steady() bool {
	return false
}
