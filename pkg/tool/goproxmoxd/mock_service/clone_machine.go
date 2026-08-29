package mock_service

import "github.com/luthermonson/go-proxmox"

func (s *Service) CloneMachine(
	instance string,
	identifier int,
	_ string,
	options *proxmox.VirtualMachineCloneOptions,
) (int, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return 0, clientFail
	}

	newIdentifier := options.NewID

	if newIdentifier == 0 {
		v, e := c.NextIdentifier()

		if e != nil {
			return 0, e
		}

		newIdentifier = v
	}

	return newIdentifier, nil
}
