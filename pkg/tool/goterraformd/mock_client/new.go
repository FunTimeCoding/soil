package mock_client

func New() *Client {
	return &Client{pods: map[string]bool{}}
}
