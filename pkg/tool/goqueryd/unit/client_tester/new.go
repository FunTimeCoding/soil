package client_tester

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(
	t *testing.T,
	address string,
) *client.Client {
	t.Helper()
	result := client.New(
		locator.NewOrigin(address),
		generativeConstant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result
}
