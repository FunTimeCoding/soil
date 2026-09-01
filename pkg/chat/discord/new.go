package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
)

func New(token string) *Client {
	client, e := discordgo.New(fmt.Sprintf("Bot %s", token))
	errors.PanicOnError(e)

	if false {
		console.Format("Identify: %+v\n", client.Identify)
		console.Format("State: %+v\n", client.State)
	}

	return &Client{client: client}
}
