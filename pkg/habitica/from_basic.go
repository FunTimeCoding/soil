package habitica

import "github.com/funtimecoding/soil/pkg/habitica/basic"

func FromBasic(b *basic.Client) *Client {
	return &Client{basic: b}
}
