package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()

	return NewWithHold(t, constant.ChannelHold)
}
