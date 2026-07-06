package secure_shell

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func KnownHosts() ssh.HostKeyCallback {
	result, e := knownhosts.New(
		join.Absolute(system.Home(), ConfigurationDirectory, KnownHostsFile),
	)
	errors.PanicOnError(e)

	return result
}
