package client_tester

func New(host string) *Client {
	return &Client{Host: host}
}
