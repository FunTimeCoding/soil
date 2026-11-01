package face

import "context"

type Connection interface {
	Open(x context.Context) error
	Serve(x context.Context) error
	Close()
}
