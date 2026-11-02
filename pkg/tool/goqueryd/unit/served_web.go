package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func servedWeb(
	t *testing.T,
) (*model_context_server.Server, *client.ClientWithResponses) {
	t.Helper()
	s := served(
		t,
		service_tester.NewOption(t, newEmbedder(t), newReranker(t)),
		memory.New(),
	)

	return s, restClient(t, s)
}
