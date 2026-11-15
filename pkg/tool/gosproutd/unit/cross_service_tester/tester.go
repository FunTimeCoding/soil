package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"testing"
)

type Tester struct {
	Service     *service.Service
	Coordinator *client.Client
	Goclauded   *base.Server
	clock       *clock_tester.Clock
	t           *testing.T
}
