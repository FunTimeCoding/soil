package telemetry

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
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

	return &Client{client: client.New(l, "").WithTimeout(5 * time.Second)}
}
