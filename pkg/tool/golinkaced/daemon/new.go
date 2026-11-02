package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/constant"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/option"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/service"
)

func New(
	o *option.LinkAce,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	c := o.Source()
	g := peer.New(constant.LinkAcePeer, c.Address(), c, l, r)

	return &Daemon{client: c, guard: g, service: service.New(c, g)}
}
