package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service"
	"time"
)

type Tester struct {
	Reporter *memory.Memory
	revealed []string
	now      func() time.Time
	service  *service.Service
}
