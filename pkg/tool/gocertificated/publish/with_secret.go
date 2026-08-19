package publish

func (p *Publisher) WithSecret(
	authority string,
	path string,
) *Publisher {
	p.secretAuthority = authority
	p.secretPath = path

	return p
}
