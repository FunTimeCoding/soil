package entry

import "maragu.dev/gomponents"

func (e *Entry) Beside(n ...gomponents.Node) *Entry {
	e.beside = append(e.beside, n...)

	return e
}
