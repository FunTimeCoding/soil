package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"testing"
)

type Tester struct {
	Service *service.Service
	t       *testing.T
}
