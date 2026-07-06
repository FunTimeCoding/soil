package mock_service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) CloneMachine(
	c face.ProxmoxClient,
	identifier int,
	_ string,
	options *proxmox.VirtualMachineCloneOptions,
) (int, error) {
	newID := options.NewID

	if newID == 0 {
		v, e := c.NextIdentifier()

		if e != nil {
			return 0, e
		}

		newID = v
	}

	return newID, nil
}
