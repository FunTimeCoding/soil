package stall_client

func (b *Builder) Verbose() *Builder {
	b.verbose = true

	return b
}
