//go:build browser

package connect

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"testing"
)

func TestAClientBuiltBeforeTheBrowserConnectsOnFirstNeed(t *testing.T) {
	s := base.New(t)
	s.Stop()
	c := chromium.New("localhost", s.Port)
	defer c.Close()
	_, e := c.Targets()
	assert.String(t, "unreachable", classify.Class(e))
	s.Start()
	_, f := c.Targets()
	assert.FatalOnError(t, f)
}

func TestAClientComesBackAfterTheBrowserRestarts(t *testing.T) {
	s := base.New(t)
	_, e := s.Client.Targets()
	assert.FatalOnError(t, e)
	s.Stop()
	_, f := s.Client.Targets()
	assert.String(t, "unreachable", classify.Class(f))
	s.Start()
	_, g := s.Client.Targets()
	assert.FatalOnError(t, g)
}
