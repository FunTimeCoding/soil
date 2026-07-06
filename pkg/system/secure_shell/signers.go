package secure_shell

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io"
)

func Signers(w io.ReadWriter) []ssh.Signer {
	result, e := agent.NewClient(w).Signers()
	errors.PanicOnError(e)

	return result
}
