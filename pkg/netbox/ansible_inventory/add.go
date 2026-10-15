package ansible_inventory

func (i *Inventory) add(
	group string,
	host string,
	variables map[string]any,
) {
	if i.groups[group] == nil {
		i.groups[group] = map[string]map[string]any{}
	}

	i.groups[group][host] = variables
}
