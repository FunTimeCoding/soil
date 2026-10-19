package mock_service

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

func New(
	instanceName string,
	proxmoxClient face.ProxmoxClient,
	snippetClient face.SnippetClient,
) *Service {
	return &Service{
		inventory: inventory.New[instance.Instance](
			instance.Instance{Name: instanceName, Host: "mock"},
		),
		selection:     selection.New(),
		proxmoxClient: proxmoxClient,
		snippetClient: snippetClient,
	}
}
