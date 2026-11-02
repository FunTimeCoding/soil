package create_machine

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/luthermonson/go-proxmox"
	"strings"
)

func (m *Machine) BuildOptions() []proxmox.VirtualMachineOption {
	var result []proxmox.VirtualMachineOption
	result = append(result, option(constant.NameOption, m.Name))
	cores := m.Cores

	if cores == 0 {
		cores = 2
	}

	result = append(result, option(constant.CoresOption, cores))
	result = append(result, option(constant.SocketsOption, 1))
	memory := m.Memory

	if memory == 0 {
		memory = 2048
	}

	result = append(result, option(constant.MemoryOption, memory))
	result = append(
		result,
		option(constant.DiskControllerOption, "virtio-scsi-pci"),
	)
	diskStorage := m.DiskStorage

	if diskStorage == "" {
		diskStorage = "local-lvm"
	}

	result = append(result, m.diskImportOptions(diskStorage)...)
	result = append(result, option(constant.BootOption, "order=virtio0;net0"))
	result = append(result, option(constant.BalloonOption, 0))
	bridge := m.Bridge

	if bridge == "" {
		bridge = "vmbr0"
	}

	model := "virtio"

	if m.HardwareAddress != "" {
		model = fmt.Sprintf("virtio=%s", m.HardwareAddress)
	}

	result = append(
		result,
		option(
			constant.PrimaryNetworkOption,
			fmt.Sprintf("%s,bridge=%s", model, bridge),
		),
	)
	agent := m.Agent == nil || *m.Agent

	if agent {
		result = append(result, option(constant.AgentOption, 1))
	}

	if m.OnBoot != nil {
		if *m.OnBoot {
			result = append(result, option(constant.OnBootOption, 1))
		} else {
			result = append(result, option(constant.OnBootOption, 0))
		}
	}

	cpuType := m.CPUType

	if cpuType == "" {
		cpuType = "host"
	}

	result = append(result, option(constant.ProcessorOption, cpuType))

	if m.OperatingSystemType != "" {
		result = append(
			result,
			option(constant.OperatingSystemOption, m.OperatingSystemType),
		)
	}

	if m.CIUser != "" {
		result = append(result, option(constant.CloudInitUserOption, m.CIUser))
	}

	if m.CIPassword != "" {
		result = append(
			result,
			option(constant.CloudInitPasswordOption, m.CIPassword),
		)
	}

	if m.SSHKeys != "" {
		keys := strings.Split(m.SSHKeys, "\n")
		result = append(
			result,
			option(
				constant.SecureShellKeysOption,
				proxmox.EncodeSSHKeys(
					keys...,
				),
			),
		)
	}

	result = append(result, m.cloudInitOptions(diskStorage)...)

	if m.Tags != "" {
		result = append(result, option(constant.TagsOption, m.Tags))
	}

	extras := parseExtras(m.Extras)
	result = append(result, extras...)

	return result
}
