package jellyfin

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/basic"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	port int,
	token string,
) *Client {
	return &Client{
		basic: basic.New(locator.New(host).Port(port).Insecure(), token),
	}
}
