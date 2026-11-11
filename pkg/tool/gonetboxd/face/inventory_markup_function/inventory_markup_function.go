package inventory_markup_function

import "context"

type InventoryMarkupFunction func(
	x context.Context,
	field string,
	site string,
) (string, error)
