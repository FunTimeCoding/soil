package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/outage_server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestToolsOnAlertmanagerOutage(t *testing.T) {
	outage_tester.New(t).Tools(
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.AlertmanagerPeer)
		},
		webConstant.OutageReaches,
		nil,
		nil,
	)
}

func TestToolsOnPrometheusOutage(t *testing.T) {
	outage_tester.New(t).Tools(
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.PrometheusPeer)
		},
		webConstant.OutageElsewhere,
		map[string]webConstant.OutageReach{
			constant.ListAlerts:    webConstant.OutageReaches,
			constant.CreateSilence: webConstant.OutageReaches,
		},
		nil,
	)
}
