package secure_shell

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"golang.org/x/crypto/ssh"
)

func Dial(
	address string,
	c *ssh.ClientConfig,
) *ssh.Client {
	result, e := ssh.Dial(constant.Transmission, address, c)
	errors.PanicOnError(e)

	return result
}
