package mock_client

func New() *Client {
	return &Client{pages: map[string]*entry{}, nextIdentifier: 1}
}
