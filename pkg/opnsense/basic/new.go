package basic

func New(
	host string,
	key string,
	secret string,
	insecure bool,
) *Client {
	return &Client{
		host:     host,
		key:      key,
		secret:   secret,
		insecure: insecure,
	}
}
