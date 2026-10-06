package xref

func (i *Index) Directory(unit string) string {
	return i.directories[unit]
}
