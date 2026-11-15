package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/unit/outage_server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestOperationsOnInstanceOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.Operations(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r)
		},
		spec,
		webConstant.OutageReaches,
		map[string]webConstant.OutageReach{
			"ListInstances": webConstant.OutageElsewhere,
			"ListBackups":   webConstant.OutageElsewhere,
			"DeleteBackup":  webConstant.OutageElsewhere,
			"CreateBackup":  webConstant.OutageElsewhere,
		},
		map[string]string{"name": "2030-01-02T03-04-05"},
		nil,
	)
}

func TestToolsOnInstanceOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r)
		},
		webConstant.OutageReaches,
		map[string]webConstant.OutageReach{
			constant.ListBackups:  webConstant.OutageElsewhere,
			constant.DeleteBackup: webConstant.OutageElsewhere,
			constant.CreateBackup: webConstant.OutageElsewhere,
		},
		map[string]map[string]any{
			constant.RestoreBackup: {
				constant.BackupParameter: "2030-01-02T03-04-05",
			},
		},
	)
}
