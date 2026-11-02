package face

import "context"

type InventoryMarkup interface {
	AnsibleInventoryMarkup(
		x context.Context,
		field string,
	) (string, error)
}
