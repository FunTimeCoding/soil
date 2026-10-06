package xref

func (i *Index) Count() int {
	return len(i.directories)
}
