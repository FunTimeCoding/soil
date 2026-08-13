package mock_client

func New() *Client {
	return &Client{groups: map[int]*group{}, nextID: 1}
}
