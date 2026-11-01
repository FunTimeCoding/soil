package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/base"
	"testing"
)

type Tester struct {
	QueryClient  *model_context_client.Client
	MemoryClient *model_context_client.Client
	QueryServer  *base.Server
	t            *testing.T
}
