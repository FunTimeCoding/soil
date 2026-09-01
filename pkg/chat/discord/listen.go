package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system"
)

func (c *Client) Listen(block bool) {
	c.client.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	c.client.AddHandler(c.receive)
	c.Open()

	if block {
		defer c.Close()
	}

	console.Line("Running")

	if block {
		system.KillSignalBlock()
	}
}
