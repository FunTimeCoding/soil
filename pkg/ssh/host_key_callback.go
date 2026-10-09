package ssh

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func (c *Client) hostKeyCallback() (ssh.HostKeyCallback, error) {
	if !c.secure {
		return ssh.InsecureIgnoreHostKey(), nil
	}

	result, e := knownhosts.New(
		join.Absolute(
			system.Home(),
			constant.SecureShellConfigurationDirectory,
			constant.KnownHostsFile,
		),
	)

	if e != nil {
		return nil, fmt.Errorf("read known hosts: %w", e)
	}

	return result, nil
}
