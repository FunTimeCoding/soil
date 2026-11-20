package check_tester

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
)

func (o *Tester) InventoryRun(
	root string,
	index string,
	i *inventory.Inventory,
) *check_result.Result {
	o.t.Helper()
	p := option.New()
	p.Root = root
	p.Index = index
	p.Inventory = i

	return o.Run(p)
}
