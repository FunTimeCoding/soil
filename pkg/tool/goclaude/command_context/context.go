package command_context

import "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"

type Context struct {
	host   string
	port   int
	client *client.ClientWithResponses
}
