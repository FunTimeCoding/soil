package web_client

import "github.com/funtimecoding/soil/pkg/face"

func New(c face.Clock) *Client {
	return &Client{clock: c}
}
