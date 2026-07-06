package firefox

import "github.com/funtimecoding/soil/pkg/firefox/constant"

func New(o ...Option) *Client {
	result := &Client{
		address: constant.DefaultHost,
		pending: make(map[int]chan *reply),
	}

	for _, f := range o {
		f(result)
	}

	return result
}
