package option

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
)

type Source struct {
	Address       string
	ServiceTokens []string
	Inventory     *inventory.Inventory
	Memory        func() face.Client
}
