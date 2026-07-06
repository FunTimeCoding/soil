package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
)

type Tester struct {
	Service  *service.Service
	Notifier *mock_notifier.Notifier
}
