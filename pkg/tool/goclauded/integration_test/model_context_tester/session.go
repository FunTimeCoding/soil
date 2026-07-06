package model_context_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"testing"
)

type Session struct {
	*model_context_client.Client
	T          *testing.T
	Context    context.Context
	RestClient *client.ClientWithResponses
	UUID       string
	PoolName   string
}
