package telegram

import "fmt"

func (c *Client) PrintDatabase() {
	if c.store == nil {
		return
	}

	f := Format
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
