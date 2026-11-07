package option

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Source struct {
	option.Web
	Inventory      *inventory.Inventory
	IndexDirectory string
	Memory         func() face.Client
}
