package godashboardd

import (
	"github.com/funtimecoding/soil/pkg/argocd"
	argocdConstant "github.com/funtimecoding/soil/pkg/argocd/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
)

func argocdClient(b *board.Board) *argocd.Client {
	if !b.HasWidget(constant.ArgocdWidget) {
		return nil
	}

	target := b.Connection.Argocd

	return argocd.New(
		target.Host,
		board.Port(target),
		target.Secure,
		environment.Required(argocdConstant.TokenEnvironment),
	)
}
