package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/unit/base"
	"testing"
)

func restClient(t *testing.T) *client.ClientWithResponses {
	t.Helper()

	return restClientOf(t, base.New(t))
}
