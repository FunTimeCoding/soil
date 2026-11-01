package pager

func (p *Pager) Next(
	path string,
	label string,
) *Pager {
	p.nextPath = path
	p.nextLabel = label

	return p
}
