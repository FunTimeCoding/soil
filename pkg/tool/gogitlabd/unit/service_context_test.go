package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
)

func TestCallerWhoLeftTakesForgeCallWithIt(t *testing.T) {
	s, clients := service_tester.New(t, constant.FixtureInstance)
	clients[constant.FixtureInstance].SeedProject(1, "group", "app")
	x, cancel := context.WithCancel(t.Context())
	cancel()
	_, e := s.Projects(x, constant.FixtureInstance, "")
	assert.True(t, errors.Is(e, context.Canceled))
}
