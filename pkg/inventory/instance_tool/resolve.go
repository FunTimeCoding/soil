package instance_tool

import "context"

func (t *Tool[T, P]) Resolve(x context.Context) (string, error) {
	return t.inventory.ResolveInstance(t.ActiveName(x))
}
