package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/option"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/service"
)

func New(
	o *option.Habitica,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	c := o.Source()
	g := peer.New(constant.HabiticaPeer, o.Host, c, l, r)

	return &Daemon{client: c, guard: g, service: service.New(c, g)}
}
