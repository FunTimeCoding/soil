package fact

func (f *Fact) Columns(count int) *Fact {
	f.columns = count

	return f
}
