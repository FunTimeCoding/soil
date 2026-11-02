package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"testing"
)

func guard(
	t *testing.T,
	name string,
	address string,
	c face.Reachable,
	l *logger.Logger,
	r face.Reporter,
) *peer.Guard {
	t.Helper()
	result := peer.New(name, address, c, l, r)
	t.Cleanup(result.Close)

	return result
}
