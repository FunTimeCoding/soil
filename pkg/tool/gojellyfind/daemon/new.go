package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/constant"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/option"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/service"
)

func New(
	o *option.Jellyfin,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	c := o.Source()
	g := peer.New(constant.JellyfinPeer, c.Address(), c, l, r)

	return &Daemon{client: c, guard: g, service: service.New(c, g)}
}
