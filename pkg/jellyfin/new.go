package jellyfin

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/transport"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func New(
	host string,
	port int,
	token string,
) *Client {
	return &Client{
		base: locator.New(host).Port(port).Insecure().String(),
		http: &http.Client{Transport: transport.New(token)},
	}
}
