package inventory

func (i *Inventory) Names() []string {
	var result []string

	for _, v := range i.Instances {
		result = append(result, v.Name)
	}

	return result
}
