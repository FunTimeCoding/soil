package loki

import "github.com/funtimecoding/soil/pkg/prometheus/loki/basic"

func New(
	host string,
	user string,
	password string,
	verbose bool,
) *Client {
	return &Client{basic: basic.New(host, user, password, verbose)}
}
