package firefox

import (
	"github.com/funtimecoding/soil/pkg/firefox/constant"
	"github.com/funtimecoding/soil/pkg/firefox/types/message"
)

func New(o ...Option) *Client {
	result := &Client{
		address: constant.DefaultHost,
		pending: make(map[int]chan *message.Reply),
	}

	for _, f := range o {
		f(result)
	}

	return result
}
