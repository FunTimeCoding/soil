package commit

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	memoryFace "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn"
)

func New(
	i *inventory.Inventory,
	w *workspace.Service,
	m memoryFace.Client,
	g *peer.Guard,
	r face.Reporter,
) *Service {
	return &Service{
		inventory: i,
		workspace: w,
		memory:    m,
		guard:     g,
		reporter:  r,
		turns:     turn.New(constant.TurnBound),
	}
}
