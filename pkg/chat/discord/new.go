package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/funtimecoding/soil/pkg/errors"
)

func New(token string) *Client {
	client, e := discordgo.New(fmt.Sprintf("Bot %s", token))
	errors.PanicOnError(e)

	return &Client{client: client}
}
