package adapter

import "context"

func (a *Adapter) Serve(x context.Context) error {
	return a.serve(x)
}
