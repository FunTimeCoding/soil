package mock_service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
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

	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return 0, clientFail
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
