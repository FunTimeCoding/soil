package inventory

func (i *Inventory[T, P]) Find(name string) (*T, bool) {
	for k := range i.Instances {
		if P(&i.Instances[k]).InstanceName() == name {
			return &i.Instances[k], true
		}
	}

	return nil, false
}
