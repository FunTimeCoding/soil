package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	"testing"
)

type Tester struct {
	*Services
	Directory string
	Memory    *mock_client.Client
	Reporter  *memory.Memory
	t         *testing.T
}
