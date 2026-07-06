package secure_shell

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func ParsePrivateKey(pemBytes []byte) ssh.Signer {
	result, e := ssh.ParsePrivateKey(pemBytes)
	errors.PanicOnError(e)

	return result
}
