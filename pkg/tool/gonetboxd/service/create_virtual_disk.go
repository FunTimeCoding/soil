package service

import "github.com/funtimecoding/soil/pkg/netbox/virtual_disk"

func (s *Service) CreateVirtualDisk(
	machine string,
	name string,
	size int32,
) (*virtual_disk.Disk, error) {
	return s.client.CreateVirtualDisk(machine, name, size)
}
