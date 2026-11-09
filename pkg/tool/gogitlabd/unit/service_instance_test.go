package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
)

func TestServiceRefusesUnknownInstance(t *testing.T) {
	s, _ := service_tester.New(t, constant.FixtureInstance)
	_, e := s.Projects(t.Context(), "charlie", "")
	assert.True(t, not_found.Is(e))
	assert.String(t, "instance not found: charlie", e.Error())
}
