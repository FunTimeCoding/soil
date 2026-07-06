//go:build local

package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestResolveCallerDatabaseError(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	s.Store.Store.Close()
	result := a.MustCallToolError(
		constant.Status,
		nil,
	)
	assert.StringContains(t, "unexpected error", result)
}
