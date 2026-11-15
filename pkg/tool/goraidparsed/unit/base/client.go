package base

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/client"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func Client(
	t *testing.T,
	p face.TopStatsParser,
) *client.Client {
	t.Helper()
	v := New(t, p)
	result := client.New(
		locator.New("localhost").Port(v.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result
}
