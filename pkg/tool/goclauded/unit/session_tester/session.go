package session_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"testing"
)

type Session struct {
	*model_context_client.Client
	t               *testing.T
	Context         context.Context
	GeneratedClient *client.ClientWithResponses
	UUID            string
	PoolName        string
}
