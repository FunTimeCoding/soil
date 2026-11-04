package stall_client

func (b *Builder) Traced() *Builder {
	b.traced = true

	return b
}
