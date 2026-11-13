package fact

func (f *Fact) Glance() *Fact {
	f.glance = true

	return f
}
