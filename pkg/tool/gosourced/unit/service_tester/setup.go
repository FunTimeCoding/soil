package service_tester

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
)

type setup struct {
	fixture    string
	directory  string
	modLines   []string
	memoryPeer face.Client
	inventory  *inventory.Inventory
}
