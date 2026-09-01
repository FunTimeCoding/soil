package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func (c *Client) receive(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case constant.DiscordPingCommand:
		c.Send(s, c.UserChannel(s, m.Author.ID).ID, "pong")
	case constant.DiscordCleanCommand:
		if m.GuildID == "" {
			// Direct message
			channel := c.UserChannel(s, m.Author.ID).ID
			c.Clean(s, channel, true)
			c.Send(s, channel, "Done")
		} else {
			// Channel
			c.Clean(s, m.ChannelID, false)
			c.Send(s, m.ChannelID, "Done")
		}
	case constant.DiscordDetailsCommand:
		console.Format("Channel: %+v\n", c.Channel(s, m.ChannelID))
	default:
		console.Format("Message: %+v\n", m.Message)
	}
}
