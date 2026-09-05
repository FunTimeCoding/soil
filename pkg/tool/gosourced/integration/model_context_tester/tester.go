package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/base"
)

type Tester struct {
	server    *base.Server
	Client    *model_context_client.Client
	Directory string
}
