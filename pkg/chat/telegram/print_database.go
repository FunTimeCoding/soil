package telegram

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/telegram/constant"
)

func (c *Client) PrintDatabase() {
	f := constant.Format
	fmt.Println("Channels:")

	for _, h := range c.channels {
		fmt.Println(h.Format(f))
	}

	fmt.Println()
	fmt.Println("Users:")

	for _, u := range c.users {
		fmt.Println(u.Format(f))
	}
}
