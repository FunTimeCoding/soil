package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/integration_test/base"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

type Tester struct {
	server *base.Server
	Client *model_context_client.Client
	Store  *store.Store
}
