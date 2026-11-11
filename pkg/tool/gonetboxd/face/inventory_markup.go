package face

import "context"

type InventoryMarkup interface {
	AnsibleInventoryMarkup(
		x context.Context,
		field string,
		site string,
	) (string, error)
}
