package embedder_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"testing"
)

type Tester struct {
	service *service.Service
	t       *testing.T
}
