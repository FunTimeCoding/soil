package connect

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
)

func Wait(l *logger.Logger) face.Client {
	result := client.NewEnvironment()
	errors.PanicOnError(
		peer.Await(
			context.Background(),
			l,
			constant.Identity.Name(),
			result.Address(),
			result.Reachable,
		),
	)

	return result
}
