package unit

import (
	"github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/unit/service_tester"
	"testing"
)

func TestGuard(t *testing.T) {
	stream := mock_stream.New(nil, false, nil)
	t.Cleanup(stream.Close)
	s := base.New(t, service_tester.Fixed(stream)).Served()
	s.VerifyBase(t)
	s.VerifyGuarded(t, "/api/status")
	s.VerifyGuarded(t, "/api/run")
}
