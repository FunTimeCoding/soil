package publish_tester

import (
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"testing"
)

type Tester struct {
	Service *service.Service
	Forge   *mock_client.Client
	t       *testing.T
}
