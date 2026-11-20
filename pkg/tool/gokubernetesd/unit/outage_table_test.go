package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/outage_server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestToolsOnClusterOutage(t *testing.T) {
	pods := map[string]any{"resourceType": "pods", "name": "alfa"}
	local, l := system.ClaimPort()
	errors.PanicClose(l)
	outage_tester.New(t).Tools(
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r)
		},
		webConstant.OutageReaches,
		map[string]webConstant.OutageReach{
			constant.ListClusters:    webConstant.OutageElsewhere,
			constant.UseCluster:      webConstant.OutageElsewhere,
			constant.ListMutedEvents: webConstant.OutageElsewhere,
			constant.MuteEvent:       webConstant.OutageElsewhere,
			constant.UnmuteEvent:     webConstant.OutageElsewhere,
			constant.StopPortForward: webConstant.OutageElsewhere,
		},
		map[string]map[string]any{
			constant.Get:      pods,
			constant.Describe: pods,
			constant.Delete:   pods,
			constant.Patch: {
				"resourceType": "pods",
				"name":         "alfa",
				"patch":        "{}",
			},
			constant.RolloutRestart: {
				"resourceType": "deployments",
				"name":         "alfa",
			},
			constant.Top: {"resourceType": "pods"},
			constant.PortForward: {
				"resourceType": "pod",
				"name":         "alfa",
				"localPort":    local,
				"targetPort":   80,
			},
			constant.ExecInPod: {"timeout": 5000},
			constant.Apply: {
				"manifest": "apiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: alfa\n",
			},
		},
	)
}
