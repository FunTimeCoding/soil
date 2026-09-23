package server

func (b *Builder) WithConnected(v func()) *Builder {
	b.connected = v

	return b
}
