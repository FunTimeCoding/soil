package service

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

type Service struct {
	inventory *instance.Inventory
	pool      *pool.Pool[instance.Instance, *instance.Instance, face.ProxmoxClient]
	snippets  *pool.Pool[instance.Instance, *instance.Instance, face.SnippetClient]
	selection *selection.Selection
}
