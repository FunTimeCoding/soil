package base

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/client"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/daemon_tester"
	"testing"
)

func Index(
	t *testing.T,
	directory string,
) *client.Client {
	t.Helper()
	d, _ := daemon_tester.NewInIndex(t, inventory.New(), directory)

	return connect(t, d, generativeConstant.ModelContextTestToken)
}
