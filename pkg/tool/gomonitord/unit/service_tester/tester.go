package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/service"
	"testing"
)

type Tester struct {
	service *service.Service
	t       *testing.T
}
