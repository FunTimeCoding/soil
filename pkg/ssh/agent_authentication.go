package ssh

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/ssh/constant"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"os"
)

func agentAuthentication() (ssh.AuthMethod, error) {
	socket, e := net.Dial("unix", os.Getenv(constant.SocketEnvironment))

	if e != nil {
		return nil, fmt.Errorf("dial agent: %w", e)
	}

	defer errors.LogClose(socket)
	signers, f := agent.NewClient(socket).Signers()

	if f != nil {
		return nil, fmt.Errorf("list agent keys: %w", f)
	}

	return ssh.PublicKeys(signers...), nil
}
