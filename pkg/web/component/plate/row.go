package plate

import "maragu.dev/gomponents"

func (p *Plate) Row(
	label string,
	value string,
) *Plate {
	return p.RowNode(label, gomponents.Text(value))
}
