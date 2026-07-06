package secure_shell

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"golang.org/x/crypto/ssh"
	"net"
)

func Listen(
	c *ssh.Client,
	address string,
) net.Listener {
	result, e := c.Listen(constant.Transmission, address)
	errors.PanicOnError(e)

	return result
}
