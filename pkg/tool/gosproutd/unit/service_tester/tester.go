package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"testing"
	"time"
)

type Tester struct {
	Service  *service.Service
	Notifier *mock_notifier.Notifier
	now      *time.Time
	t        *testing.T
}
