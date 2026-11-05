package client_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/opnsense/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"testing"
)

func New(
	t *testing.T,
	s face.OpnsenseClient,
) *client.Client {
	t.Helper()

	return NewWith(t, s, memory.New())
}
