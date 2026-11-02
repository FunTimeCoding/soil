package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	memoryConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/option"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
)

func New(
	o *option.Source,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	m := o.Memory()
	g := peer.New(memoryConstant.Identity.Name(), m.Address(), m, l, r)

	return &Daemon{
		memory:  m,
		guard:   g,
		service: service.New(o.Inventory, o.IndexDirectory, m, g, r),
	}
}
