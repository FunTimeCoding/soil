package band

import (
	"crypto/tls"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func New(
	address string,
	user string,
	password string,
) *Client {
	errors.FatalOnEmpty(address, "address")
	errors.FatalOnEmpty(user, "user")
	errors.FatalOnEmpty(password, "password")

	return &Client{
		base:     address,
		user:     user,
		password: password,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}
