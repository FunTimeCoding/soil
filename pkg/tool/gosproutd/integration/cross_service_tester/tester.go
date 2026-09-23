package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"testing"
	"time"
)

type Tester struct {
	Service     *service.Service
	Coordinator *connector.Client
	Goclauded   *base.Server
	now         *time.Time
	t           *testing.T
}
