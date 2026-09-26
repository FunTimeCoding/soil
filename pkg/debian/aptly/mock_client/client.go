package mock_client

type Client struct {
	versions map[string][]string
	fail     error
}
