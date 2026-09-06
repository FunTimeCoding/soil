package inventory

func New(instances ...Instance) *Inventory {
	return &Inventory{Instances: instances}
}
