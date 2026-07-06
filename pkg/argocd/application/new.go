package application

import "github.com/funtimecoding/soil/pkg/argocd/response"

func New(v response.Application) *Application {
	return &Application{
		Name:   v.Metadata.Name,
		Sync:   v.Status.Sync.Status,
		Health: v.Status.Health.Status,
	}
}
