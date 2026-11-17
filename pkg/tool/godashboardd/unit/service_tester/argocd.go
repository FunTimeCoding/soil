package service_tester

import "github.com/funtimecoding/soil/pkg/kubernetes/face"

func (o *Tester) argocd() face.ArgocdClient {
	if o.Argocd == nil {
		return nil
	}

	return o.Argocd
}
