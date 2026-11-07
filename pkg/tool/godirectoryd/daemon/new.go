package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/web"
)

func New(
	o *option.Directory,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	d := o.Source()
	g := peer.New(constant.DirectoryPeer, o.Host, d, l, r)
	s := service.New(d, g)

	return &Daemon{guard: g, service: s, web: web.New(s)}
}
