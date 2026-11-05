package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/option"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/service"
)

func New(
	o *option.Sublime,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	c := o.Source()
	g := peer.New(constant.SublimePeer, c.Address(), c, l, r)

	return &Daemon{client: c, guard: g, service: service.New(c, g)}
}
