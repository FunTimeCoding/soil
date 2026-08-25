package mock_client

func New() *Client {
	return &Client{groups: map[int]*group{}, nextIdentifier: 1}
}
