package inventory

func (i *Inventory[T, P]) Names() []string {
	result := make([]string, 0, len(i.Instances))

	for k := range i.Instances {
		result = append(result, P(&i.Instances[k]).InstanceName())
	}

	return result
}
