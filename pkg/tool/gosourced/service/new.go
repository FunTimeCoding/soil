package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	lintFace "github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/source/index/cache"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	memoryFace "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn"
)

func New(
	i *inventory.Inventory,
	indexDirectory string,
	m memoryFace.Client,
	g *peer.Guard,
	r face.Reporter,
) *Service {
	return &Service{
		inventory:  i,
		workspaces: cache.New(indexDirectory, lintFace.Kind()),
		memory:     m,
		guard:      g,
		reporter:   r,
		turns:      turn.New(constant.TurnBound),
	}
}
