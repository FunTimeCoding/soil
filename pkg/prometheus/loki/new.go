package loki

import "github.com/funtimecoding/soil/pkg/prometheus/loki/basic"

// Conflicts with github.com/prometheus/common/config
//import "github.com/grafana/loki/v3/integration/client"

func New(
	host string,
	user string,
	password string,
	verbose bool,
) *Client {
	//client.New()
	return &Client{basic: basic.New(host, user, password, verbose)}
}
