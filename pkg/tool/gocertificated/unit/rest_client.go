package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/base"
	"github.com/funtimecoding/soil/pkg/web"
	"testing"
)

func restClient(
	t *testing.T,
	s *base.Server,
) *client.ClientWithResponses {
	t.Helper()
	result, e := client.NewClientWithResponses(
		s.Address(""),
		client.WithHTTPClient(s.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)

	return result
}
