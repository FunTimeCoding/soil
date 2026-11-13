package toolchain

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
	"testing"
)

func inventoryRun(
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

	return checked(t, o)
}
