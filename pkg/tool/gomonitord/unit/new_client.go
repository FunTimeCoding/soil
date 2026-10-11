package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"testing"
)

func newClient(
	t *testing.T,
	port int,
) *client.ClientWithResponses {
	t.Helper()
	result, e := client.NewClientWithResponses(
		fmt.Sprintf("http://localhost:%d", port),
		client.WithRequestEditorFn(
			web.BearerEditor(constant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)

	return result
}
