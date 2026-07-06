package iterm

import (
	"github.com/funtimecoding/soil/pkg/iterm/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

func New(o ...Option) *Client {
	result := &Client{
		base:   locator.New(constant.DefaultHost).Insecure(),
		client: &http.Client{},
	}

	for _, f := range o {
		f(result)
	}

	return result
}
