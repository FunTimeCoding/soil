package counting_source

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"

func New(s face.OpnsenseSource) *Client {
	return &Client{OpnsenseSource: s}
}
