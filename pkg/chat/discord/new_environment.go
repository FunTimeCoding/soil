package discord

import "github.com/funtimecoding/soil/pkg/system/environment"

func NewEnvironment() *Client {
	return New(environment.Required(TokenEnvironment))
}
