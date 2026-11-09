package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"testing"
)

func TestPreCompactStartsCompaction(t *testing.T) {
	phase, known := goclaude.CompactionPhase(constant.PreCompactEvent)
	assert.True(t, known)
	assert.String(t, "started", string(phase))
}

func TestPostCompactFinishesCompaction(t *testing.T) {
	phase, known := goclaude.CompactionPhase(constant.PostCompactEvent)
	assert.True(t, known)
	assert.String(t, "finished", string(phase))
}

func TestAnotherHookEventIsNoCompaction(t *testing.T) {
	_, known := goclaude.CompactionPhase("SessionStart")
	assert.False(t, known)
}
