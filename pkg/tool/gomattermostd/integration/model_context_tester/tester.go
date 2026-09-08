package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/integration/base"
)

type Tester struct {
	Client *model_context_client.Client
	Server *base.Server
}
