package mock_client

func New() *Client {
	return &Client{nextIdentifier: 1}
}
