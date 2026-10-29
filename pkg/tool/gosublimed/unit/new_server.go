package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generated "github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/server"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/unit/service_tester"
	"testing"
)

func newServer(
	t *testing.T,
	content string,
) *server.Server {
	t.Helper()
	s := server.New(service_tester.New(t, mock_client.New()), memory.New())
	_, e := s.CreateView(
		context.Background(),
		generated.CreateViewRequestObject{
			Body: &generated.CreateViewJSONRequestBody{
				Title:   "notes",
				Content: content,
			},
		},
	)

	if e != nil {
		t.Fatal(e)
	}

	return s
}
