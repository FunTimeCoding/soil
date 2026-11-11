package inventory_markup_function

import "context"

func (f InventoryMarkupFunction) AnsibleInventoryMarkup(
	x context.Context,
	field string,
	site string,
) (string, error) {
	return f(x, field, site)
}
