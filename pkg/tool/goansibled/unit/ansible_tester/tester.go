package ansible_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/service"
	"testing"
)

type Tester struct {
	t        *testing.T
	Service  *service.Service
	Reporter *memory.Memory
	capture  string
	litePath string
}
