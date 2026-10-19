package instance_tool

import "context"

func (t *Tool[T, P]) Resolve(x context.Context) (string, error) {
	name, e := t.ActiveName(x)

	if e != nil {
		return "", e
	}

	return t.inventory.ResolveInstance(name)
}
