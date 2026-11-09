package chromium

import "github.com/funtimecoding/soil/pkg/web"

func NewCombined(combined string) *Client {
	host, port := web.LocatorHostPortSplit(combined)

	return New(host, port)
}
