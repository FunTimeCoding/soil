package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	microsoftConstant "github.com/funtimecoding/soil/pkg/microsoft/constant"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gographd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gographd/option"
	"github.com/funtimecoding/soil/pkg/tool/gographd/service"
)

func New(
	o *option.Server,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	c := o.Source()
	g := peer.New(constant.GraphPeer, microsoftConstant.GraphHost, c, l, r)

	return &Daemon{client: c, guard: g, service: service.New(c, g)}
}
