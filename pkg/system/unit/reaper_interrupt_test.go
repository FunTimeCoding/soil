package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"testing"
)

func TestReaperInterruptSkipsVanishedProcess(t *testing.T) {
	m := memory.New()
	r := reaper.New(m)
	defer r.Close()
	r.Add(1 << 30)
	r.Interrupt()
	assert.Count(t, 0, m.Events())
}
