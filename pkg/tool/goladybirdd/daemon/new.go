package daemon

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/navigation"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/option"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/service"
)

func New(
	o *option.Ladybird,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	w := navigation.New(constant.NavigationWindow, constant.NavigationLimit)
	c := o.Client(w.Observe)
	errors.PanicOnError(c.Subscribe(navigation.Events()...))
	g := peer.New(constant.LadybirdPeer, c.Address(), c, l, r)

	return &Daemon{
		client:  c,
		guard:   g,
		service: service.New(c, g, w, o.DownloadDirectory),
	}
}
