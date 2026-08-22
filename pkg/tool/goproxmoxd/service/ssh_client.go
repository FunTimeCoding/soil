package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
)

func (s *Service) SSHClient(instance string) (face.SnippetClient, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if c, okay := s.sshClients[instance]; okay {
		return c, nil
	}

	i, okay := s.Instance(instance)

	if !okay {
		return nil, not_found.New("instance", instance)
	}

	if i.SSHUser == "" || i.SSHPassword == "" {
		return nil, not_configured.Format(
			"instance %s has no SSH credentials configured",
			instance,
		)
	}

	c := ssh.NewWithPassword(i.SSHUser, i.Host, i.SSHPassword, false)
	s.sshClients[instance] = c

	return c, nil
}
