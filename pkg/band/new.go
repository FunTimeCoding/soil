package band

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(
	address string,
	user string,
	password string,
) *Client {
	errors.FatalOnEmpty(address, "address")
	errors.FatalOnEmpty(user, "user")
	errors.FatalOnEmpty(password, "password")
	p := stall_client.New().Insecure().Build()

	return &Client{
		base:     address,
		user:     user,
		password: password,
		client:   p.Client(),
		pool:     p,
	}
}
