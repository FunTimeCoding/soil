package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/double"
	"testing"
)

func NewWithSessionFile(
	t *testing.T,
	identifier string,
	slug string,
) *Server {
	t.Helper()
	d := double.New(t)
	t.Cleanup(d.Close)
	d.WriteSessionFile(identifier, slug)

	return serve(t, constant.ChannelHold, d)
}
