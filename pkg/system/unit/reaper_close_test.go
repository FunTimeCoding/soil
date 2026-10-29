package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"testing"
)

func TestReaperClosesTwice(t *testing.T) {
	r := reaper.New(memory.New())
	r.Start()
	r.Close()
	r.Close()
}
