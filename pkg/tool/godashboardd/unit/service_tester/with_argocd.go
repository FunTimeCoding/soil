package service_tester

import "github.com/funtimecoding/soil/pkg/kubernetes/argocd/mock_client"

func WithArgocd() Option {
	return func(o *Tester) {
		o.Argocd = mock_client.New(o.t)
	}
}
