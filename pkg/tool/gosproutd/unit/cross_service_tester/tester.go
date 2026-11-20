package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
)

type Tester struct {
	Service     *service.Service
	Coordinator *client.Client
	Goclauded   *base.Server
	store       *store_tester.Tester
	t           *testing.T
}
