package fact

import "maragu.dev/gomponents"

func (f *Fact) Row(
	label string,
	value string,
) *Fact {
	return f.RowNode(label, gomponents.Text(value))
}
