package basic

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func New(
	host string,
	port int,
	user string,
	password string,
	eauth string,
) *Client {
	result := &Client{
		base:     locator.New(host).Port(port).Insecure().String(),
		user:     user,
		password: password,
		eauth:    eauth,
		client:   &http.Client{},
	}
	result.login()

	return result
}
