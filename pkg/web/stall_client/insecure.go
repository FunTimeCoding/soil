package stall_client

func (b *Builder) Insecure() *Builder {
	b.insecure = true

	return b
}
