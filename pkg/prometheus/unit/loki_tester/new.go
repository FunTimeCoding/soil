package loki_tester

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()
	result := basic.New(
		locator.NewOrigin(server).Base(constant.LokiBase),
		"alfa",
		"bravo-password",
		false,
	)
	t.Cleanup(result.Close)

	return result
}
