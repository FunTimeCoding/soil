package option

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
)

type Gitlab struct {
	Address       string
	ServiceTokens []string
	MetricAddress string
	Inventory     *instance.Inventory
	Client        func(*instance.Instance) (face.GitLabClient, error)
}
