package discord

import "github.com/funtimecoding/soil/pkg/console"

func (c *Client) PrintStatus() {
	console.Format("User: %+v\n", c.User().Username)

	for _, guild := range c.Guilds() {
		console.Format("Guild: %+v\n", guild.ID)

		for _, channel := range c.Channels(guild.ID) {
			console.Format("Channel: %s\n", channel.Name)
		}
	}
}
