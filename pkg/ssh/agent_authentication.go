package ssh

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/ssh/constant"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io"
	"net"
	"os"
)

func agentAuthentication() (ssh.AuthMethod, io.Closer, error) {
	socket, e := net.Dial("unix", os.Getenv(constant.SocketEnvironment))

	if e != nil {
		return nil, nil, fmt.Errorf("dial agent: %w", e)
	}

	signers, f := agent.NewClient(socket).Signers()

	if f != nil {
		errors.LogClose(socket)

		return nil, nil, fmt.Errorf("list agent keys: %w", f)
	}

	return ssh.PublicKeys(signers...), socket, nil
}
