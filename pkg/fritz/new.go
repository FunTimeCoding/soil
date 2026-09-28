package fritz

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/fritz/constant"
	"net/http"
)

func New(
	host string,
	user string,
	password string,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(user, "user")
	errors.FatalOnEmpty(password, "password")

	return &Client{
		base:     fmt.Sprintf("http://%s:%d", host, constant.Port),
		user:     user,
		password: password,
		client:   &http.Client{},
	}
}
