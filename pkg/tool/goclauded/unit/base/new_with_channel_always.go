package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/peer_tester"
	"testing"
)

func NewWithChannelAlways(t *testing.T) *Server {
	t.Helper()
	d := peer_tester.New(t)
	d.ChannelAlways = true
	t.Cleanup(d.Close)

	return serve(t, constant.ChannelHold, d)
}
