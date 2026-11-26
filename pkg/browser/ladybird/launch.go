package ladybird

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/run"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net"
	"strconv"
)

func (c *Client) launch() {
	port, l := system.ClaimPort()
	errors.PanicClose(l)
	c.address = net.JoinHostPort(webConstant.Loopback, strconv.Itoa(port))
	c.driver = run.New().Open(c.arguments(port)...)
}
