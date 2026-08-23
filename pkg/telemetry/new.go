package telemetry

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
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

	return &Client{
		base:   l.String(),
		client: &http.Client{Timeout: 5 * time.Second},
	}
}
