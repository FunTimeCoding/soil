package inventory

func (i *Inventory[T, P]) ResolveInstance(explicit string) (string, error) {
	return Resolve(explicit, i.Names())
}
