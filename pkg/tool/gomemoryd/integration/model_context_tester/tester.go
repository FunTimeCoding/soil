package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/base"
	"testing"
)

type Tester struct {
	*model_context_client.Client
	t    *testing.T
	base *base.Server
}
