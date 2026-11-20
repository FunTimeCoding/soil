package service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
)

type Tester struct {
	Service  *service.Service
	Notifier *mock_notifier.Notifier
	store    *store_tester.Tester
	t        *testing.T
}
