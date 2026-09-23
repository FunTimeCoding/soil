package server

func (b *Builder) WithExperimental(v map[string]any) *Builder {
	b.experimental = v

	return b
}
