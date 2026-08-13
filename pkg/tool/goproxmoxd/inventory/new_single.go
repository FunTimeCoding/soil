package inventory

func NewSingle(name string) *Inventory {
	return &Inventory{Instances: []Instance{{Name: name, Host: "mock"}}}
}
