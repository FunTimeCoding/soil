package service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/mock_notifier"
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"testing"
)

type Tester struct {
	Service  *service.Service
	Notifier *mock_notifier.Notifier
	clock    *clock_tester.Clock
	t        *testing.T
}
