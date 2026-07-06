package basic

import (
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	user string,
	password string,
	verbose bool,
) *Client {
	// https://grafana.com/docs/loki/latest/reference/loki-http-api
	return &Client{
		user:     user,
		password: password,
		verbose:  verbose,
		base:     locator.New(host).Base(constant.Base),
	}
}
