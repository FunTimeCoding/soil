package mock_client

func New() *Client {
	return &Client{versions: make(map[string][]string)}
}
