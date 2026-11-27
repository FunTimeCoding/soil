package firefox

import "github.com/funtimecoding/soil/pkg/browser/firefox/types/message"

func New() *Client {
	return &Client{pending: make(map[int]chan *message.Reply)}
}
