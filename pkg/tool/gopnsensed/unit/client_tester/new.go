package client_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"testing"
)

func New(
	t *testing.T,
	s face.OpnsenseSource,
) *client.Client {
	t.Helper()

	return NewWith(t, s, memory.New())
}
