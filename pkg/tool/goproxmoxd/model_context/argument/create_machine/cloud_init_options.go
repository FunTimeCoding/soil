package create_machine

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/luthermonson/go-proxmox"
)

func (m *Machine) cloudInitOptions(diskStorage string) []proxmox.VirtualMachineOption {
	var result []proxmox.VirtualMachineOption
	cloudInit := m.CIUser != "" || m.SSHKeys != "" || m.CIPassword != ""

	if cloudInit {
		ipConfiguration := m.IPConfiguration

		if ipConfiguration == "" {
			ipConfiguration = "ip=dhcp"
		}

		result = append(
			result,
			option(constant.InternetConfigurationOption, ipConfiguration),
		)

		if m.SearchDomain != "" {
			result = append(
				result,
				option(constant.SearchDomainOption, m.SearchDomain),
			)
		}

		result = append(
			result,
			option(
				constant.RemovableDriveOption,
				fmt.Sprintf("%s:cloudinit", diskStorage),
			),
		)
	} else if m.CDROM != "" {
		result = append(
			result,
			option(
				constant.RemovableDriveOption,
				fmt.Sprintf("%s,media=cdrom", m.CDROM),
			),
		)
	}

	return result
}
