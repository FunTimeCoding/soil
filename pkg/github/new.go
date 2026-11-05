package github

func New(token string) *Client {
	return build(token)
}
