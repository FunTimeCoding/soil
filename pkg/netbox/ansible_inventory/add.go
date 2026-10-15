package ansible_inventory

func (i *Inventory) add(
	group string,
	host string,
	address string,
) {
	if i.groups[group] == nil {
		i.groups[group] = map[string]string{}
	}

	i.groups[group][host] = address
}
