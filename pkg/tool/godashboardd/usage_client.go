package godashboardd

import (
	nextcloudConstant "github.com/funtimecoding/soil/pkg/nextcloud/constant"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
)

func usageClient(b *board.Board) *usage.Client {
	if !b.HasWidget(constant.NextcloudWidget) {
		return nil
	}

	target := b.Connection.Nextcloud

	return usage.New(
		target.Host,
		board.Port(target),
		target.Secure,
		environment.Required(nextcloudConstant.TokenEnvironment),
	)
}
