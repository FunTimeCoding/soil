package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/proxmox/network_device"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
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

	if m.HardwareAddress == "" {
		address, e := network_device.Derive(
			s.inventory.Index(instance),
			identifier,
		)

		if e != nil {
			return 0, e
		}

		m.HardwareAddress = address
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
			constant.PrimaryDiskOption,
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
