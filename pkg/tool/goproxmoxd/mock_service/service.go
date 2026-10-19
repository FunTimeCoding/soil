package mock_service

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

type Service struct {
	inventory     *instance.Inventory
	selection     *selection.Selection
	proxmoxClient face.ProxmoxClient
	snippetClient face.SnippetClient
}
