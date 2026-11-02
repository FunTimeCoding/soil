package create_machine

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/luthermonson/go-proxmox"
)

func (m *Machine) diskImportOptions(diskStorage string) []proxmox.VirtualMachineOption {
	var result []proxmox.VirtualMachineOption

	if m.DiskImport != "" {
		result = append(
			result,
			option(
				constant.PrimaryDiskOption,
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
				constant.PrimaryDiskOption,
				fmt.Sprintf(
					"%s:%d,aio=io_uring,backup=1,cache=none,discard=on,iothread=1,replicate=1",
					diskStorage,
					diskSize,
				),
			),
		)
	}

	return result
}
