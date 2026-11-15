package service_tester

import (
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	bandMock "github.com/funtimecoding/soil/pkg/band/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/service"
	gonetboxdMock "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"testing"
)

func Intent(t *testing.T) (
	*service.Service,
	*gonetboxdMock.Client,
	*bandMock.Client,
	*memory.Memory,
) {
	t.Helper()
	r := memory.New()
	v, engines, n := NewWithNetbox(t, r, "alfa")
	engine := engines["alfa"]
	engine.SetPower(bandConstant.PowerOn, RunningStates()...)

	return v, n, engine, r
}
