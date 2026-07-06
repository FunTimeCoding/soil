package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/integration_test/base"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
)

type Tester struct {
	server     *base.Server
	Client     *model_context_client.Client
	MockClient *mock_client.Client
}
