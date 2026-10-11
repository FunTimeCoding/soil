package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/unit/base"
	"testing"
)

type Tester struct {
	server *base.Server
	Client *model_context_client.Client
	Store  *store.Store
	t      *testing.T
}
