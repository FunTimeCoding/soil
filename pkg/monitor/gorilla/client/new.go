package client

import (
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/location"
	"net/url"
)

func New() *Client {
	return &Client{
		locator: url.URL{
			Scheme: constant.Socket,
			Host: web.AddressHostPort(
				constant.Localhost,
				constant.ListenPort,
			),
			Path: location.Monitor,
		},
		done: make(chan struct{}),
	}
}
