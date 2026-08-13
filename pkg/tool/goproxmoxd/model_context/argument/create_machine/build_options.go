package create_machine

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
	"strings"
)

func (m *Machine) BuildOptions() []proxmox.VirtualMachineOption {
	var result []proxmox.VirtualMachineOption
	result = append(result, option("name", m.Name))
	cores := m.Cores

	if cores == 0 {
		cores = 2
	}

	result = append(result, option("cores", cores))
	result = append(result, option("sockets", 1))
	memory := m.Memory

	if memory == 0 {
		memory = 2048
	}

	result = append(result, option("memory", memory))
	result = append(result, option("scsihw", "virtio-scsi-pci"))
	diskStorage := m.DiskStorage

	if diskStorage == "" {
		diskStorage = "local-lvm"
	}

	if m.DiskImport != "" {
		result = append(
			result,
			option(
				"virtio0",
				fmt.Sprintf(
					"%s:0,import-from=%s,aio=io_uring,backup=1,cache=none,discard=on,iothread=1,replicate=1",
					diskStorage,
					m.DiskImport,
				),
			),
		)
	} else {
		diskSize := m.DiskSize

		if diskSize == 0 {
			diskSize = 32
		}

		result = append(
			result,
			option(
				"virtio0",
				fmt.Sprintf(
					"%s:%d,aio=io_uring,backup=1,cache=none,discard=on,iothread=1,replicate=1",
					diskStorage,
					diskSize,
				),
			),
		)
	}

	result = append(result, option("boot", "order=virtio0;net0"))
	result = append(result, option("balloon", 0))
	bridge := m.Bridge

	if bridge == "" {
		bridge = "vmbr0"
	}

	result = append(
		result,
		option("net0", fmt.Sprintf("virtio,bridge=%s", bridge)),
	)
	agent := m.Agent == nil || *m.Agent

	if agent {
		result = append(result, option("agent", 1))
	}

	if m.OnBoot != nil {
		if *m.OnBoot {
			result = append(result, option("onboot", 1))
		} else {
			result = append(result, option("onboot", 0))
		}
	}

	cpuType := m.CPUType

	if cpuType == "" {
		cpuType = "host"
	}

	result = append(result, option("cpu", cpuType))

	if m.OSType != "" {
		result = append(result, option("ostype", m.OSType))
	}

	if m.CIUser != "" {
		result = append(result, option("ciuser", m.CIUser))
	}

	if m.CIPassword != "" {
		result = append(result, option("cipassword", m.CIPassword))
	}

	if m.SSHKeys != "" {
		keys := strings.Split(m.SSHKeys, "\n")
		result = append(
			result,
			option("sshkeys", proxmox.EncodeSSHKeys(keys...)),
		)
	}

	cloudInit := m.CIUser != "" || m.SSHKeys != "" || m.CIPassword != ""

	if cloudInit {
		ipConfiguration := m.IPConfiguration

		if ipConfiguration == "" {
			ipConfiguration = "ip=dhcp"
		}

		result = append(result, option("ipconfig0", ipConfiguration))

		if m.SearchDomain != "" {
			result = append(result, option("searchdomain", m.SearchDomain))
		}

		result = append(
			result,
			option("ide2", fmt.Sprintf("%s:cloudinit", diskStorage)),
		)
	} else if m.CDROM != "" {
		result = append(
			result,
			option("ide2", fmt.Sprintf("%s,media=cdrom", m.CDROM)),
		)
	}

	if m.Tags != "" {
		result = append(result, option("tags", m.Tags))
	}

	extras := parseExtras(m.Extras)
	result = append(result, extras...)

	return result
}
