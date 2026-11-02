package unit

import (
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/service"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/service_tester"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"testing"
)

func intentService(t *testing.T) (
	*service.Service,
	*netbox.Client,
	*mock_client.Client,
	*memory.Memory,
) {
	t.Helper()
	r := memory.New()
	v, engines, n := service_tester.NewWithNetbox(t, r, "alfa")
	engine := engines["alfa"]
	engine.SetPower(bandConstant.PowerOn, runningStates()...)

	return v, n, engine, r
}
