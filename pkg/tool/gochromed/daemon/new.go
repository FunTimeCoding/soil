package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/option"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/service"
)

func New(
	o *option.Chrome,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	c := o.Source()
	g := peer.New(constant.ChromiumPeer, c.Address(), c, l, r)

	return &Daemon{
		source:  c,
		guard:   g,
		service: service.New(c, g, o.DownloadDirectory),
	}
}
