package fact

func (f *Fact) Tooltip(text string) *Fact {
	if len(f.rows) == 0 {
		return f
	}

	f.rows[len(f.rows)-1].Tooltip = text

	return f
}
