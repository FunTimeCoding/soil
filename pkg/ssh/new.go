package ssh

import "github.com/funtimecoding/soil/pkg/ssh/constant"

func New(
	user string,
	host string,
	secure bool,
) *Client {
	return &Client{
		user:         user,
		host:         host,
		port:         constant.Port,
		secure:       secure,
		authenticate: agentAuthentication,
		Panic:        true,
	}
}
