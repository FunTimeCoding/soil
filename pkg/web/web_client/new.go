package web_client

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(c face.Clock) *Client {
	return &Client{clock: c, client: stall_client.New().Build().Client()}
}
