package mock_service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"
)

func New(
	instanceName string,
	proxmoxClient face.ProxmoxClient,
	snippetClient face.SnippetClient,
) *Service {
	return &Service{
		inventory:      inventory.NewSingle(instanceName),
		proxmoxClient:  proxmoxClient,
		snippetClient:  snippetClient,
		activeInstance: make(map[string]string),
	}
}
