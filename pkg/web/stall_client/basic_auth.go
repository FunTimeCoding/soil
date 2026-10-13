package stall_client

func (b *Builder) BasicAuth(
	user string,
	password string,
) *Builder {
	b.basic = true
	b.user = user
	b.password = password

	return b
}
