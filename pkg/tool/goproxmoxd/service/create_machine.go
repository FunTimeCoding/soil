package service

import (
	"fmt"
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
	task, e := c.CreateMachine(node, identifier, options...)

	if e != nil {
		return 0, e
	}

	e = c.WaitForTask(task, 300)

	if e != nil {
		return 0, e
	}

	if m.DiskImport != "" {
		diskSize := m.DiskSize

		if diskSize == 0 {
			diskSize = 32
		}

		vm, e := c.Machine(node, identifier)

		if e != nil {
			return 0, e
		}

		resizeTask, e := c.ResizeDisk(
			vm,
			"virtio0",
			fmt.Sprintf("%dG", diskSize),
		)

		if e != nil {
			return 0, e
		}

		e = c.WaitForTask(resizeTask, 120)

		if e != nil {
			return 0, e
		}
	}

	if m.Start {
		vm, e := c.Machine(node, identifier)

		if e != nil {
			return 0, e
		}

		startTask, e := c.StartMachine(vm)

		if e != nil {
			return 0, e
		}

		e = c.WaitForTask(startTask, 120)

		if e != nil {
			return 0, e
		}
	}

	return identifier, nil
}
