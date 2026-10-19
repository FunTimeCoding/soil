package instance_tool

import "context"

func (t *Tool[T, P]) ActiveName(x context.Context) (string, error) {
	return t.selection.Active(x)
}
