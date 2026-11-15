package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/peer_tester"
	"testing"
	"time"
)

func NewWithHold(
	t *testing.T,
	hold time.Duration,
) *Server {
	t.Helper()
	d := peer_tester.New(t)
	t.Cleanup(d.Close)

	return serve(t, hold, d)
}
