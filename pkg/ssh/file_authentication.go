package ssh

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"golang.org/x/crypto/ssh"
	"os"
)

func fileAuthentication(
	keyPath string,
	keyName string,
) (ssh.AuthMethod, error) {
	root, e := os.OpenRoot(keyPath)

	if e != nil {
		return nil, fmt.Errorf("open key directory: %w", e)
	}

	defer errors.LogClose(root)
	b, f := root.ReadFile(keyName)

	if f != nil {
		return nil, fmt.Errorf("read key: %w", f)
	}

	signer, g := ssh.ParsePrivateKey(b)

	if g != nil {
		return nil, fmt.Errorf("parse key: %w", g)
	}

	return ssh.PublicKeys(signer), nil
}
