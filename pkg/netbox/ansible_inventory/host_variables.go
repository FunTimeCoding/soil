package ansible_inventory

import "github.com/funtimecoding/soil/pkg/netbox/constant"

func hostVariables(
	primary string,
	tunnel string,
	configuration map[string]any,
) map[string]any {
	result := map[string]any{}

	if tunnel != "" {
		result[constant.InventoryAddress] = tunnel
	} else if a := hostAddress(primary); a != "" {
		result[constant.InventoryAddress] = a
	}

	if c, okay := configuration[constant.InventoryContext].(map[string]any); okay {
		for k, v := range c {
			result[k] = v
		}
	}

	return result
}
