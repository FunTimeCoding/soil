package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument/create_machine"
)

func (s *Service) CreateMachine(
	instance string,
	m *create_machine.Machine,
) (int, error) {
	cloudInit := m.CIUser != "" || m.SSHKeys != "" || m.CIPassword != ""

	if cloudInit && m.CDROM != "" {
		return 0, validation.New(
			"cdrom and cloud-init are mutually exclusive - both use ide2",
		)
	}

	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (int, error) {
			return s.provisionMachine(c, instance, m)
		},
	)
}
