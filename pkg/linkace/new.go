package linkace

import (
	"github.com/funtimecoding/soil/pkg/linkace/basic"
	"github.com/funtimecoding/soil/pkg/linkace/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	token string,
) *Client {
	return &Client{
		basic: basic.New(locator.New(host).Base(constant.BasePath), token),
		host:  host,
	}
}
