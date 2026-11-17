package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/peer_tester"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Server {
	t.Helper()
	p := &setup{hold: constant.ChannelHold}

	for _, f := range o {
		f(p)
	}

	d := peer_tester.New(t)
	d.ChannelAlways = p.channelAlways

	for _, f := range p.sessionFiles {
		d.WriteSessionFile(f[0], f[1])
	}

	return serve(t, p.hold, d)
}
