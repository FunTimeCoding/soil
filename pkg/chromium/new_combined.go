package chromium

import "github.com/funtimecoding/soil/pkg/web"

func NewCombined(combined string) *Client {
	host, port := web.HostPortFromLocatorSplit(combined)

	return New(host, port)
}
