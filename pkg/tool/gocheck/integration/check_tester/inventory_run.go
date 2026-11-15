package check_tester

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
	"testing"
)

func InventoryRun(
	t *testing.T,
	root string,
	index string,
	i *inventory.Inventory,
) *check_result.Result {
	t.Helper()
	o := option.New()
	o.Root = root
	o.Index = index
	o.Inventory = i

	return Run(t, o)
}
