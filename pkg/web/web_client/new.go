package web_client

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/web"
)

func New(c face.Clock) *Client {
	return &Client{clock: c, client: web.StallClient()}
}
