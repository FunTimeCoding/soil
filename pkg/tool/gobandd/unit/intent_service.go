package unit

import (
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/service"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"testing"
)

func intentService(t *testing.T) (
	*service.Service,
	*netbox.Client,
	*mock_client.Client,
	*memory.Memory,
) {
	t.Helper()
	engine := mock_client.New()
	engine.SetPower(constant.PowerOn, runningStates()...)
	n := netbox.New()
	r := memory.New()
	o := alfaOption()
	o.Netbox = func() face.Client { return n }
	v := service.New(o, logger.New(t.Context()), r)
	t.Cleanup(v.Close)
	v.Pool().Set("alfa", engine)

	return v, n, engine, r
}
