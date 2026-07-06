package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/funtimecoding/soil/pkg/system"
)

func (c *Client) Listen(block bool) {
	c.client.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	c.client.AddHandler(c.receive)
	c.Open()

	if block {
		defer c.Close()
	}

	fmt.Println("Running")

	if block {
		system.KillSignalBlock()
	}
}
