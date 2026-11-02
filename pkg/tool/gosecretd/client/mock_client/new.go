package mock_client

func New() *Client {
	return &Client{Secrets: map[string]map[string]string{}}
}
