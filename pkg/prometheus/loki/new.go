package loki

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	user string,
	password string,
	verbose bool,
) *Client {
	return &Client{
		basic: basic.New(
			locator.New(host).Base(constant.LokiBase),
			user,
			password,
			verbose,
		),
	}
}
