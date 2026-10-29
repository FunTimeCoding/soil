package godashboardd

import (
	nextcloudConstant "github.com/funtimecoding/soil/pkg/nextcloud/constant"
	"github.com/funtimecoding/soil/pkg/nextcloud/face"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage/client"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board/connection"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
)

func usageClient(b *board.Board) func() face.UsageClient {
	return func() face.UsageClient {
		if !b.HasWidget(constant.NextcloudWidget) {
			return nil
		}

		target := b.Connection.Nextcloud

		return client.New(
			target.Host,
			connection.Port(target),
			target.Secure,
			environment.Required(nextcloudConstant.TokenEnvironment),
		)
	}
}
