package service

import (
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/types/host"
	"runtime"
)

func (s *Service) Host() *host.Host {
	return host.New(
		system.Hostname(),
		runtime.GOOS,
		network.HardwareAddresses(),
	)
}
