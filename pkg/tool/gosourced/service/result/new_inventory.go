package result

func NewInventory(
	region string,
	total int,
	more int,
	calls []*Call,
) *Inventory {
	return &Inventory{Region: region, Total: total, More: more, Calls: calls}
}
