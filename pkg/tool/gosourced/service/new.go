package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	lintFace "github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/index/cache"
	memoryConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/option"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/turn"
)

func New(
	o *option.Source,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	m := o.Memory()

	return &Service{
		inventory:  o.Inventory,
		workspaces: cache.New(index.DefaultDirectory(), lintFace.Kind()),
		memory:     m,
		guard: peer.New(
			memoryConstant.Identity.Name(),
			m.Address(),
			m,
			l,
			r,
		),
		reporter: r,
		turns:    turn.New(constant.TurnBound),
	}
}
