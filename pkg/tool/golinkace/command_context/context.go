package command_context

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
)

type Context struct {
	host     string
	client   *client.ClientWithResponses
	terminal *terminal.Terminal
	pool     *pool.Pool
}
