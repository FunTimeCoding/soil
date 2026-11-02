package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/directory"
	directoryConstant "github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"testing"
)

func stack(t *testing.T) *service.Service {
	t.Helper()
	d := directory.NewEnvironment()
	g := peer.New(
		constant.DirectoryPeer,
		environment.Required(directoryConstant.HostEnvironment),
		d,
		logger.New(context.Background()),
		memory.New(),
	)
	t.Cleanup(g.Close)

	return service.New(d, g)
}
