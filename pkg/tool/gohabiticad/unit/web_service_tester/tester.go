package web_service_tester

import (
	"github.com/funtimecoding/soil/pkg/habitica/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/base"
	"testing"
)

type Tester struct {
	server     *base.Server
	Client     *client.ClientWithResponses
	MockClient *mock_client.Client
	t          *testing.T
}
