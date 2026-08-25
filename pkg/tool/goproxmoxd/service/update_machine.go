package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument/update_machine"
	"github.com/luthermonson/go-proxmox"
	"strings"
)

func (s *Service) UpdateMachine(
	c face.ProxmoxClient,
	a *update_machine.Machine,
) error {
	if e := a.Validate(); e != nil {
		return e
	}

	var options []proxmox.VirtualMachineOption

	if a.Name != "" {
		options = append(
			options,
			proxmox.VirtualMachineOption{
				Name:  constant.NameOption,
				Value: a.Name,
			},
		)
	}

	if a.Tags != "" {
		options = append(
			options,
			proxmox.VirtualMachineOption{
				Name:  constant.TagsOption,
				Value: a.Tags,
			},
		)
	}

	if a.OnBoot != nil {
		value := 0

		if *a.OnBoot {
			value = 1
		}

		options = append(
			options,
			proxmox.VirtualMachineOption{
				Name:  constant.OnBootOption,
				Value: value,
			},
		)
	}

	if a.Cores > 0 {
		options = append(
			options,
			proxmox.VirtualMachineOption{
				Name:  constant.CoresOption,
				Value: a.Cores,
			},
		)
	}

	if a.Memory > 0 {
		options = append(
			options,
			proxmox.VirtualMachineOption{
				Name:  constant.MemoryOption,
				Value: a.Memory,
			},
		)
	}

	if a.Description != "" {
		options = append(
			options,
			proxmox.VirtualMachineOption{
				Name:  constant.DescriptionOption,
				Value: a.Description,
			},
		)
	}

	if a.Delete != "" {
		options = append(
			options,
			proxmox.VirtualMachineOption{
				Name:  constant.DeleteOption,
				Value: strings.TrimSpace(a.Delete),
			},
		)
	}

	return s.UpdateMachineConfiguration(c, a.Identifier, a.Node, options)
}
