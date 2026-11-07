package option

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Gitlab struct {
	option.Web
	Metric    option.Web
	Inventory *instance.Inventory
	Client    func(*instance.Instance) (face.GitLabClient, error)
}
