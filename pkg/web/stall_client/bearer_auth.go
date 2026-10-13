package stall_client

func (b *Builder) BearerAuth(token string) *Builder {
	b.bearer = true
	b.token = token

	return b
}
