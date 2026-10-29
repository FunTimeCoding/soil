package godashboardd

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/argocd"
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board/connection"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
)

func argocdClient(b *board.Board) func() face.ArgocdClient {
	return func() face.ArgocdClient {
		if !b.HasWidget(constant.ArgocdWidget) {
			return nil
		}

		target := b.Connection.Argocd

		return argocd.New(
			target.Host,
			connection.Port(target),
			target.Secure,
			environment.Required(constant.ArgocdTokenEnvironment),
		)
	}
}
