package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gobandd/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory"
)

func New(i *inventory.Inventory) *Service {
	return &Service{inventory: i, engines: make(map[string]face.Engine)}
}
