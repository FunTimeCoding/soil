package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/integration_test/base"
)

type Tester struct {
	Server *base.Server
	Client *model_context_client.Client
}
