package unit

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newLokiClient(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()

	return basic.New(
		upstream_tester.Locator(t, server).Base(constant.LokiBase),
		"alfa",
		"bravo-password",
		false,
	)
}
