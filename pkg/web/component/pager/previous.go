package pager

func (p *Pager) Previous(
	path string,
	label string,
) *Pager {
	p.previousPath = path
	p.previousLabel = label

	return p
}
