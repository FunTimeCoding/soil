package godashboardd

import (
	"github.com/funtimecoding/soil/pkg/argocd"
	"github.com/funtimecoding/soil/pkg/argocd/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
)

func argocdClient(b *board.Board) *argocd.Client {
	if !b.HasWidget(board.ArgocdWidget) {
		return nil
	}

	target := b.Connection.Argocd

	return argocd.New(
		target.Host,
		board.Port(target),
		target.Secure,
		environment.Required(constant.TokenEnvironment),
	)
}
