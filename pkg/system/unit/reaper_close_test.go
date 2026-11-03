package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func TestReaperClosesTwice(t *testing.T) {
	r := reaper.New(memory.New(), prometheus.NewRegistry())
	r.Start()
	r.Close()
	r.Close()
}
