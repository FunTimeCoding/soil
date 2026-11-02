package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func TestAMomentReachesTheMemoryPeer(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.Moment,
		map[string]any{constant.Line: "the build went green"},
	)
	assert.Count(t, 1, s.MemoryPeer.Impressions)
	assert.String(
		t,
		"the build went green",
		s.MemoryPeer.Impressions[0].Content,
	)
}
