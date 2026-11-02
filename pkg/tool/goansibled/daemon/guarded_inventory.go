package daemon

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face/inventory_markup_function"
)

func guardedInventory(
	c face.Client,
	g *peer.Guard,
) inventory_markup_function.InventoryMarkupFunction {
	return func(
		x context.Context,
		field string,
	) (string, error) {
		var result string
		e := g.Attempt(
			x,
			func(y context.Context) error {
				var f error
				result, f = c.AnsibleInventoryMarkup(y, field)

				return f
			},
		)

		return result, e
	}
}
