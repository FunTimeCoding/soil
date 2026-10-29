package option

import (
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	nextcloud "github.com/funtimecoding/soil/pkg/nextcloud/face"
)

func New() *Dashboard {
	return &Dashboard{
		Argocd: func() kubernetes.ArgocdClient { return nil },
		Usage:  func() nextcloud.UsageClient { return nil },
	}
}
