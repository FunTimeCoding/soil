package instance

func Index(
	i *Inventory,
	name string,
) int {
	if v, okay := i.Find(name); okay {
		return v.Index
	}

	return 0
}
