package firefox

import "github.com/funtimecoding/soil/pkg/firefox/types/message"

func New() *Client {
	return &Client{pending: make(map[int]chan *message.Reply)}
}
