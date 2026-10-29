package mock_client

func New() *Client {
	return &Client{Logs: map[string][]string{}}
}
