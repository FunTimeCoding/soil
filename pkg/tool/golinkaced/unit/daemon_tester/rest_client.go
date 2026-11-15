package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/linkace/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func RestClient(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
) *client.ClientWithResponses {
	t.Helper()
	s := served(t, c, r)
	result, e := client.NewClientWithResponses(
		s.Address(webConstant.RootPath),
		client.WithHTTPClient(s.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)

	return result
}
