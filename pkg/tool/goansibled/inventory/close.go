package inventory

func (i *Inventory) Close() {
	i.pool.Close()
}
