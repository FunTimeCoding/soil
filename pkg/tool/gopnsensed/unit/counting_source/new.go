package counting_source

import "github.com/funtimecoding/soil/pkg/opnsense/face"

func New(s face.OpnsenseClient) *Client {
	return &Client{OpnsenseClient: s}
}
