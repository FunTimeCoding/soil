package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"testing"
)

func TestResolveCallerDatabaseError(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	s.Store.Store.Close()
	result := a.MustCallToolError(constant.SessionStatus, nil)
	assert.StringContains(t, "unexpected error", result)
}
