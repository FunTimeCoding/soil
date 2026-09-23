package fixture

import "context"

func Cancelled() context.Context {
	x, cancel := context.WithCancel(context.Background())
	cancel()

	return x
}
