package adapter

import "context"

func (a *Adapter) Open(x context.Context) error {
	return a.open(x)
}
