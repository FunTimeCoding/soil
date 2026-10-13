package sublime

import (
	"github.com/funtimecoding/soil/pkg/sublime/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(o ...Option) *Client {
	p := stall_client.New().Build()
	result := &Client{
		base:   locator.New(constant.DefaultHost).Insecure(),
		client: p.Client(),
		pool:   p,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
