package secure_shell

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func Session(c *ssh.Client) *ssh.Session {
	result, e := c.NewSession()
	errors.PanicOnError(e)

	return result
}
