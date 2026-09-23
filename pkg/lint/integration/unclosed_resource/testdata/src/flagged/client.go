package example

type Client struct{}

func newClient() *Client {
	return &Client{}
}

func (c *Client) Work() {}

func (c *Client) Close() {}
