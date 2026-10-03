package discord

import "github.com/bwmarrin/discordgo"

func directMessage(m *discordgo.MessageCreate) bool {
	return m.GuildID == ""
}
