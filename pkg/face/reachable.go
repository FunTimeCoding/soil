package face

import "context"

type Reachable interface {
	Reachable(x context.Context) error
}
