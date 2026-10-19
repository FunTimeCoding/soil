package service

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
)

type Service struct {
	inventory *instance.Inventory
	pool      *pool.Pool[instance.Instance, *instance.Instance, face.Forge]
	selection *selection.Selection
}
