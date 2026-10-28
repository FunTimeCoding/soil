package stall_client

func (b *Builder) Browser() *Builder {
	b.browser = true

	return b
}
