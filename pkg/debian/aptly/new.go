package aptly

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func New(
	host string,
	port int,
	insecure bool,
	username string,
	password string,
) *Client {
	l := locator.New(host).Port(port)

	if insecure {
		l.Insecure()
	}

	return &Client{
		Base:     l.String(),
		Username: username,
		Password: password,
		client:   &http.Client{},
	}
}
