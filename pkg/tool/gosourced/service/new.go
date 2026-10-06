package service

import (
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/index/cache"
	"github.com/funtimecoding/soil/pkg/source/inventory"
)

func New(i *inventory.Inventory) *Service {
	return &Service{
		inventory:  i,
		workspaces: cache.New(index.DefaultDirectory(), face.Kind()),
	}
}
