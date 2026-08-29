package inventory

func (i *Inventory) Index(name string) int {
	for _, v := range i.Instances {
		if v.Name == name {
			return v.Index
		}
	}

	return 0
}
