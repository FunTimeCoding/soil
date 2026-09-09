package face

import "github.com/funtimecoding/soil/pkg/argocd/application"

type ApplicationSource interface {
	Applications() ([]*application.Application, error)
}
