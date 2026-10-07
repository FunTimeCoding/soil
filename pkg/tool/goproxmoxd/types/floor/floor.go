package floor

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/guest"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/node"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/storage"
)

type Floor struct {
	Nodes    []node.Node
	Guests   []guest.Guest
	Storages []storage.Storage
}
