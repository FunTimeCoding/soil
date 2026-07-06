package mock_service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument/create_machine"
)

func (s *Service) CreateMachine(
	c face.ProxmoxClient,
	m *create_machine.Machine,
) (int, error) {
	cloudInit := m.CIUser != "" || m.SSHKeys != "" || m.CIPassword != ""

	if cloudInit && m.CDROM != "" {
		return 0, constant.ErrorCDROMCloudInitConflict
	}

	identifier := m.Identifier

	if identifier <= 0 {
		v, e := c.NextIdentifier()

		if e != nil {
			return 0, e
		}

		identifier = v
	}

	node, e := c.Node(m.Node)

	if e != nil {
		return 0, e
	}

	options := m.BuildOptions()
	_, e = c.CreateMachine(node, identifier, options...)

	if e != nil {
		return 0, e
	}

	return identifier, nil
}
