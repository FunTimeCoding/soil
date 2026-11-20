package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/outage_server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestToolsOnGitLabOutage(t *testing.T) {
	outage_tester.New(t).Tools(
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r)
		},
		webConstant.OutageReaches,
		nil,
		nil,
	)
}
