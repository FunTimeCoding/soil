package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/unit/service_tester"
	"testing"
	"time"
)

func TestMasterOutageFailsToolWithoutCapture(t *testing.T) {
	r := memory.New()
	s := model_context_server.Mount(t, service_tester.NewDownDaemon(t, r), r)
	defer s.Close()
	c := s.Connect(t)
	defer c.Close()
	assert.StringContains(
		t,
		"connection refused",
		c.MustCallToolError(constant.Keys, map[string]any{}),
	)
	assert.Count(t, 0, r.Events())
}

func TestKeySyncStaysQuietWhileMasterIsDown(t *testing.T) {
	r := memory.New()
	d := service_tester.NewDownDaemon(t, r)
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	time.Sleep(200 * time.Millisecond)

	for _, w := range workers {
		w.Close()
	}

	assert.Count(t, 0, r.Events())
}

func TestDaemonHandsOutMasterGuardFirst(t *testing.T) {
	_, okay := service_tester.NewDownDaemon(
		t,
		memory.New(),
	).Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}
