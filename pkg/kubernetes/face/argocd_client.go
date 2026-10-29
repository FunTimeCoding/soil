package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/argocd/application"
)

type ArgocdClient interface {
	Applications() ([]*application.Application, error)
	Reachable(x context.Context) error
	Close()
}
