package stall_client

func (b *Builder) Socket(path string) *Builder {
	b.socket = path

	return b
}
