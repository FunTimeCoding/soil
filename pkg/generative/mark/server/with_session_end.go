package server

func (b *Builder) WithSessionEnd(v func(session string)) *Builder {
	b.sessionEnd = v

	return b
}
