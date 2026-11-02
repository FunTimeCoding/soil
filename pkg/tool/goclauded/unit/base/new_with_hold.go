package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/double"
	"testing"
	"time"
)

func NewWithHold(
	t *testing.T,
	hold time.Duration,
) *Server {
	t.Helper()
	d := double.New(t)
	t.Cleanup(d.Close)

	return serve(t, hold, d)
}
