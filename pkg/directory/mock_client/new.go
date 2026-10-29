package mock_client

func New(base string) *Client {
	return &Client{
		base:      base,
		records:   map[string]map[string][]string{},
		passwords: map[string]string{},
		members:   map[string]bool{},
	}
}
