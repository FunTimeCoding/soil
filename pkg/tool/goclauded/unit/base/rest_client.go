package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func (s *Server) RESTClient(t *testing.T) *client.ClientWithResponses {
	t.Helper()
	c, e := client.NewClientWithResponses(
		locator.New(constant.Localhost).Insecure().Port(s.Port).String(),
		client.WithRequestEditorFn(
			web.BearerEditor(generative.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)

	return c
}
