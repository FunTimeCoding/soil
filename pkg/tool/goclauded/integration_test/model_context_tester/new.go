package model_context_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/google/uuid"
	"testing"
)

func New(
	t *testing.T,
	port int,
) *Session {
	t.Helper()
	x := context.Background()
	c := model_context_client.New(t, port)
	base := locator.New(
		constant.Localhost,
	).Insecure().Port(port).String()
	restClient, e := client.NewClientWithResponses(base)
	assert.FatalOnError(t, e)
	identifier := uuid.New().String()
	response, f := restClient.GetCheckWithResponse(
		x,
		&client.GetCheckParams{Session: identifier},
	)
	assert.FatalOnError(t, f)

	return &Session{
		Client:     c,
		T:          t,
		Context:    x,
		RestClient: restClient,
		UUID:       identifier,
		PoolName:   response.JSON200.Callsign,
	}
}
