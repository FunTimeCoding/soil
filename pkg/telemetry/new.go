package telemetry

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"time"
)

func New(
	host string,
	port int,
	insecure bool,
) *Client {
	l := locator.New(host).Port(port)

	if insecure {
		l.Insecure()
	}

	p := stall_client.New().Timeout(5 * time.Second).Build()

	return &Client{base: l.String(), client: p.Client(), pool: p}
}
