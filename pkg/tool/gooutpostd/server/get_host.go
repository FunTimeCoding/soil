package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/server"
	"runtime"
)

func (*Server) GetHost(
	_ context.Context,
	_ server.GetHostRequestObject,
) (server.GetHostResponseObject, error) {
	addresses := network.HardwareAddresses()

	return server.GetHost200JSONResponse{
		Hostname:          system.Hostname(),
		Platform:          runtime.GOOS,
		HardwareAddresses: &addresses,
	}, nil
}
