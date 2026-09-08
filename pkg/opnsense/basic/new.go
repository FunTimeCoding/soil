package basic

func New(
	host string,
	key string,
	secret string,
	untrusted bool,
) *Client {
	return &Client{
		host:      host,
		key:       key,
		secret:    secret,
		untrusted: untrusted,
	}
}
