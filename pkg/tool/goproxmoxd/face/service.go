package face

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

type Service interface {
	SnapshotService
	MachineService
	ContainerService
	NodeService
	Inventory() *instance.Inventory
	Selection() *selection.Selection
	Client(instance string) (ProxmoxClient, error)
	SSHClient(instance string) (SnippetClient, error)
}
