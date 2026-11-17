package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/iterm/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/service"
	"testing"
	"time"
)

type Tester struct {
	Source      *mock_client.Client
	Reporter    *memory.Memory
	keyInterval time.Duration
	service     *service.Service
	t           *testing.T
}
