package telemetry

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"time"
)

func New(
	host string,
	port int,
) *Client {
	return &Client{
		base:   locator.New(host).Insecure().Port(port).String(),
		client: &http.Client{Timeout: 5 * time.Second},
	}
}
