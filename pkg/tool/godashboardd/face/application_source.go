package face

import "github.com/funtimecoding/soil/pkg/kubernetes/argocd/application"

type ApplicationSource interface {
	Applications() ([]*application.Application, error)
}
