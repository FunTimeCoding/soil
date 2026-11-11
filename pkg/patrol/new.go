package patrol

func New(
	user string,
	host string,
	keyDirectory string,
	keyName string,
) *Client {
	return &Client{
		user:         user,
		host:         host,
		keyDirectory: keyDirectory,
		keyName:      keyName,
	}
}
