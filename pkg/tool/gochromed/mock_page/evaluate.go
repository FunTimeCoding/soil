package mock_page

import "github.com/funtimecoding/soil/pkg/notation"

func (p *Page) Evaluate(
	_ string,
	out any,
) error {
	if p.evaluation == nil {
		return nil
	}

	return notation.DecodeBytes(notation.Marshal(p.evaluation), out)
}
