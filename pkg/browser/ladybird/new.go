package ladybird

import (
	"github.com/funtimecoding/soil/pkg/browser/types/ladybird_option"
	"github.com/funtimecoding/soil/pkg/browser/types/message"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(o *ladybird_option.Option) *Client {
	return &Client{
		option:  o,
		pool:    stall_client.New().Build(),
		address: o.Address,
		pending: make(map[int]chan *message.Reply),
	}
}
