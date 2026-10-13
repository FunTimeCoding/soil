package stall_client

func (b *Builder) Long() *Builder {
	b.long = true

	return b
}
