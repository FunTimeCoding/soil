package xref

func (i *Index) Has(unit string) bool {
	_, okay := i.directories[unit]

	return okay
}
