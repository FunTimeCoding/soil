package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func servedClient(t *testing.T) *client.Client {
	t.Helper()
	s := served(
		t,
		service_tester.NewOption(t, newEmbedder(t), newReranker(t)),
		memory.New(),
	)
	result := client.New(
		locator.New(webConstant.Localhost).Insecure().Port(s.Port),
		constant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result
}
