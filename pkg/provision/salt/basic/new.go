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
	insecure bool,
) *Client {
	l := locator.New(host)

	if port != 0 {
		l.Port(port)
	}

	if insecure {
		l.Insecure()
	}

	result := &Client{
		base:     l.String(),
		user:     user,
		password: password,
		eauth:    eauth,
		client:   &http.Client{},
	}
	result.login()

	return result
}
