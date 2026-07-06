package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/integration_test/base"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/mock_client"
)

type Tester struct {
	server     *base.Server
	Client     *model_context_client.Client
	MockClient *mock_client.Client
}
