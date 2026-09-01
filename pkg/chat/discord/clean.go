package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/funtimecoding/soil/pkg/console"
)

func (c *Client) Clean(
	s *discordgo.Session,
	channel string,
	onlyOwn bool,
) {
	for i, m := range c.Messages(s, channel) {
		if onlyOwn && m.Author.ID != s.State.User.ID {
			continue
		}

		console.Format("Delete %d: %s\n", i, m.ID)
		c.Delete(channel, m)
	}
}
