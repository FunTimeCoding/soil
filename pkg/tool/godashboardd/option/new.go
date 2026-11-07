package option

import (
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/face"
	nextcloud "github.com/funtimecoding/soil/pkg/nextcloud/face"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/option"
)

func New() *Dashboard {
	return &Dashboard{
		Web:    option.Web{Authorization: client.NewEnvironment},
		Argocd: func() kubernetes.ArgocdClient { return nil },
		Usage:  func() nextcloud.UsageClient { return nil },
	}
}
