package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/opnsense/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/service"
	"testing"
)

type Tester struct {
	Reporter *memory.Memory
	source   face.OpnsenseClient
	service  *service.Service
	t        *testing.T
}
