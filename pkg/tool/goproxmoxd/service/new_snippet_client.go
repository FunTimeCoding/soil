package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

func newSnippetClient(i *instance.Instance) (face.SnippetClient, error) {
	if i.SSHUser == "" || i.SSHPassword == "" {
		return nil, not_configured.Format(
			"instance %s has no SSH credentials configured",
			i.Name,
		)
	}

	result := ssh.NewWithPassword(i.SSHUser, i.Host, i.SSHPassword, false)

	if e := result.Dial(); e != nil {
		return nil, e
	}

	return result, nil
}
