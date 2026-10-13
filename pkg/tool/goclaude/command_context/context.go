package command_context

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
)

type Context struct {
	host       string
	port       int
	client     *client.ClientWithResponses
	longClient *client.ClientWithResponses
	terminal   *terminal.Terminal
	pool       *pool.Pool
	longPool   *pool.Pool
}
