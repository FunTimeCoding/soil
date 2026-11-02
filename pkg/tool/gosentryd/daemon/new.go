package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/service"
)

func New(
	o *option.Server,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	c := o.Source()
	g := peer.New(constant.SentryPeer, o.Host, c, l, r)

	return &Daemon{
		client:  c,
		guard:   g,
		service: service.New(c, g, o.Organization),
	}
}
