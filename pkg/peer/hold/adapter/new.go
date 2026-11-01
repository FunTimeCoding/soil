package adapter

import "context"

func New(
	open func(context.Context) error,
	serve func(context.Context) error,
	close func(),
) *Adapter {
	return &Adapter{open: open, serve: serve, close: close}
}
