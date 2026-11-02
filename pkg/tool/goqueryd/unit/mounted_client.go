package unit

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/option"
	"testing"
)

func mountedClient(
	t *testing.T,
	o *option.Query,
	r face.Reporter,
) *client.ClientWithResponses {
	t.Helper()

	return restClient(t, served(t, o, r))
}
