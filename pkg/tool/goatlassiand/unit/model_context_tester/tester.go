package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/base"
	"testing"
)

type Tester struct {
	server     *base.Server
	Client     *model_context_client.Client
	MockClient *mock_client.Client
	t          *testing.T
}
