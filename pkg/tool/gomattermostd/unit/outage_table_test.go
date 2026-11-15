package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/outage_server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"path/filepath"
	"testing"
)

func TestOperationsOnMattermostOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.Operations(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.MattermostPeer)
		},
		spec,
		webConstant.OutageReaches,
		map[string]webConstant.OutageReach{
			"GetStatus": webConstant.OutageElsewhere,
		},
		map[string]string{"since": "2026-01-01T00:00:00Z"},
		nil,
	)
}

func TestOperationsOnRosterOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.Operations(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.RosterPeer)
		},
		spec,
		webConstant.OutageElsewhere,
		nil,
		map[string]string{"since": "2026-01-01T00:00:00Z"},
		nil,
	)
}

func TestToolsOnMattermostOutage(t *testing.T) {
	path := filepath.Join(t.TempDir(), "alfa.txt")
	system.SaveFile(path, "alfa")
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.MattermostPeer)
		},
		webConstant.OutageReaches,
		map[string]webConstant.OutageReach{
			constant.ListSubscriptions: webConstant.OutageElsewhere,
			constant.UnsubscribeThread: webConstant.OutageElsewhere,
		},
		map[string]map[string]any{
			constant.GetChannel:        {"channel_id": "alfa"},
			constant.GetChannelHistory: {"channel_id": "alfa"},
			constant.PostMessage:       {"channel_id": "alfa"},
			constant.UploadFile:        {"channel_id": "alfa", "path": path},
		},
	)
}

func TestToolsOnRosterOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.RosterPeer)
		},
		webConstant.OutageElsewhere,
		nil,
		nil,
	)
}
