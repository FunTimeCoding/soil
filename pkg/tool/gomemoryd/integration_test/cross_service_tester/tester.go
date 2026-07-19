package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration_test/base"
)

type Tester struct {
	QueryClient  *model_context_client.Client
	MemoryClient *model_context_client.Client
	MemoryStore  *store.Store
	QueryServer  *base.Server
}
