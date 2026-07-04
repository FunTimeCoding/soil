package board

func (b *Board) HasWidget(name string) bool {
	for _, v := range b.Entries() {
		if v.Widget == name {
			return true
		}
	}

	return false
}
