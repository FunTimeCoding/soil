package model_context_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/google/uuid"
	"testing"
)

func New(
	t *testing.T,
	s *model_context_server.Server,
) *Session {
	t.Helper()
	x := context.Background()
	c := s.Connect(t)
	restClient, e := client.NewClientWithResponses(
		s.Address(webConstant.RootPath),
		client.WithHTTPClient(s.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)
	identifier := uuid.New().String()
	response, f := restClient.GetCheckWithResponse(
		x,
		&client.GetCheckParams{Session: identifier},
	)
	assert.FatalOnError(t, f)
	result := &Session{
		Client:     c,
		T:          t,
		Context:    x,
		RestClient: restClient,
		UUID:       identifier,
		PoolName:   response.JSON200.Callsign,
	}
	t.Cleanup(result.Close)

	return result
}
