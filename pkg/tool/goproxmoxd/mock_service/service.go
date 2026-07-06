package mock_service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"
	"sync"
)

type Service struct {
	inventory      *inventory.Inventory
	proxmoxClient  face.ProxmoxClient
	snippetClient  face.SnippetClient
	activeInstance map[string]string
	mutex          sync.Mutex
}
