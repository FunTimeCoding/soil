package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/unit/outage_server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestToolsOnMasterOutage(t *testing.T) {
	outage_tester.New(t).Tools(
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r)
		},
		webConstant.OutageElsewhere,
		map[string]webConstant.OutageReach{
			constant.Keys:    webConstant.OutageReaches,
			constant.Minions: webConstant.OutageReaches,
		},
		nil,
	)
}
