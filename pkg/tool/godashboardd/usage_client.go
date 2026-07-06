package godashboardd

import (
	"github.com/funtimecoding/soil/pkg/nextcloud/constant"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
)

func usageClient(b *board.Board) *usage.Client {
	if !b.HasWidget(board.NextcloudWidget) {
		return nil
	}

	target := b.Connection.Nextcloud

	return usage.New(
		target.Host,
		board.Port(target),
		target.Secure,
		environment.Required(constant.TokenEnvironment),
	)
}
