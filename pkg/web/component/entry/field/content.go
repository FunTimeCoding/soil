package field

import "maragu.dev/gomponents"

func (f *Field) Content(n ...gomponents.Node) *Field {
	f.content = n

	return f
}
