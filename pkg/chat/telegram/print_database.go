package telegram

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/telegram/constant"
)

func (c *Client) PrintDatabase() {
	if c.store == nil {
		return
	}

	f := constant.Format
	fmt.Println("Channels:")

	for _, h := range c.store.MustChannels() {
		fmt.Println(h.Format(f))
	}

	fmt.Println()
	fmt.Println("Users:")

	for _, u := range c.store.MustUsers() {
		fmt.Println(u.Format(f))
	}
}
