package ssh

func New(
	user string,
	host string,
	secure bool,
) *Client {
	return &Client{
		user:         user,
		host:         host,
		secure:       secure,
		authenticate: agentAuthentication,
		Panic:        true,
	}
}
