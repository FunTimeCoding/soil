package client

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/socket"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func NewLocal(procfilePath string) *Client {
	return build(
		locator.New(constant.Identity.Name()).Insecure(),
		"",
		stall_client.New().Socket(socket.Path(procfilePath)).Build(),
	)
}
