package board

func (b *Board) Entries() []*Entry {
	var result []*Entry

	for _, c := range b.Top {
		for _, s := range c.Sections {
			result = append(result, s.Entries...)
		}
	}

	for _, s := range b.Tail.Sections {
		result = append(result, s.Entries...)
	}

	return result
}
