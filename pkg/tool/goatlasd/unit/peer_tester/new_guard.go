package peer_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"testing"
)

func NewGuard(
	t *testing.T,
	p face.Reachable,
) *peer.Guard {
	t.Helper()
	result := peer.New(
		t.Name(),
		t.Name(),
		p,
		logger.New(context.Background()),
		memory.New(),
	)
	t.Cleanup(result.Close)

	return result
}
